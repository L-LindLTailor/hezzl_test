package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_users = make(map[string]int32)
	new_users["Ivan"] = 43
	new_users["Pavel"] = 30
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())

	}
	params := &pb.GetUsers{}
	r, err := c.GetListUsers(ctx, params)
	if err != nil {
		log.Fatalf("colud not retrieve users: %v", err)
	}

	log.Printf("\nUSER LIST: \n")
	fmt.Printf("r.GetUsers(): %v\n:", r.GetList())
	del, err := c.DeleteUser(ctx, &pb.DelUser{UserID: 1})

	log.Printf("\nUSER DELETE: \n")
	fmt.Printf("Del User: %v\n:", del.UserID)
}
