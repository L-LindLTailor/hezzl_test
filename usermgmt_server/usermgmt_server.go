package main

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v4"
	"github.com/tech-with-moss/go-usermgmt-grpc/kafka_test"
	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

// connects
const (
	port = ":50051"
)

// user status
const (
	active   = 1
	inactive = 0
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	conn              *pgx.Conn
	firstUserCreation bool
	pb.UnimplementedUserManagementServer
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {

	clID := int32(rand.Intn(10000000))

	go kafka_test.StartKafka(
		clID,
		in.GetName(),
		in.GetAge(),
		time.Now().Format(time.RFC850),
		active,
	)

	createSql := `
    CREATE TABLE IS NOT EXISTS users(
    	id SERIAL PRIMARY KEY,
		clID int,
		name text,
		age int
    );`

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	createdUser := &pb.User{Name: in.GetName(), Age: in.GetAge()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), "INSERT INTO users (clID, name, age) VALUES ($1, $2)", clID, createdUser.Name, createdUser.Age)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())

	return createdUser, nil
}

func (server *UserManagementServer) GetListUsers(ctx context.Context, in *pb.GetUsers) (*pb.ListUsers, error) {

	conn, err := redis.Dial("tcp", "localhost:6379", redis.DialReadTimeout(time.Minute))
	checkErr(err)
	defer conn.Close()
	podcast := "podcast:1"
	var users []*pb.User

	reply, err := redis.Values(conn.Do("HGETALL", podcast))
	checkErr(err)
	err = redis.ScanStruct(reply, &users)
	checkErr(err)

	var userList *pb.ListUsers = &pb.ListUsers{}

	if users[0].Name != "" && users[0].Age != 0 {

		userList.List = users

		return userList, nil
	}

	rows, err := server.conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}

		userList.List = append(userList.List, &user)

		_, err = conn.Do(
			"HMSET",
			podcast,
			userList.List,
		)
	}

	return userList, nil
}

func (server *UserManagementServer) DeleteUser(ctx context.Context, in *pb.DelUser) (*pb.UserID, error) {

	var userDel *pb.UserID = &pb.UserID{}

	rows, err := server.conn.Query(context.Background(), "DELETE FROM users WHERE id ="+string(in.GetUserID()))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := pb.UserID{}
		err = rows.Scan(&user.UserID)
		if err != nil {
			return nil, err
		}
		userDel.UserID = user.UserID
	}

	return userDel, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	databaseUrl := "postgres://postgres:postgres@localhost:5432/postgres"
	var userMgmtServer *UserManagementServer = NewUserManagementServer()
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	userMgmtServer.conn = conn
	userMgmtServer.firstUserCreation = true
	if err := userMgmtServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
