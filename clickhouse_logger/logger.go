package clickhouse_logger

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func LoggerInClickHouse(message kafka.Message, userID int32, useName string, userAge int32, timeStamp string, userActive uint8) {
	connect, err := sql.Open("clickhouse_logger", "tcp://127.0.0.1:9000?username=&compress=true&debug=true")
	checkErr(err)
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS userLogger
		(
			userID      INTEGER,
			kafkaMSSG 	String,
			name        String,
			age         UInt8,
			action_time DateTime,
			active      UInt8
		) engine = Memory
	`)

	checkErr(err)
	tx, err := connect.Begin()
	checkErr(err)
	stmt, err := tx.Prepare("INSERT INTO userLogger (userID, message, name, age, action_time, active) VALUES (?, ?, ?, ?, ?)")
	checkErr(err)

	if _, err := stmt.Exec(
		userID,
		string(message.Value),
		useName,
		userAge,
		timeStamp,
		userActive,
	); err != nil {
		log.Fatal(err)
	}

	checkErr(tx.Commit())
	rows, err := connect.Query("SELECT userID, message, name, age, action_time, active FROM userLogger")
	checkErr(err)
	for rows.Next() {
		var (
			userID_, age_   int32
			message_, name_ string
			actionTime      time.Time
			active_         uint8
		)
		checkErr(rows.Scan(&userID_, &message_, &name_, &age_, &actionTime, &active_))
		log.Printf("userID: %d, message: %s, name: %s, age: %v, actionTime: %s, active: %d", userID_, message_, name_, age_, actionTime, active_)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
