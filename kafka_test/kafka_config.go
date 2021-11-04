package kafka_test

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/tech-with-moss/go-usermgmt-grpc/clickhouse_logger"
	"log"
)

func StartKafka(userID int32, useName string, userAge int32, timeStamp string, userActive uint8)  {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "myTopic",
		GroupID: "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		mssg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("kafka_test reader failed: %v", err)
			continue
		}
		clickhouse_logger.LoggerInClickHouse(mssg, userID, useName, userAge , timeStamp, userActive)
	}
}