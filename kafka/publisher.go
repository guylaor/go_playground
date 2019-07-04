package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func writeSome() {

	//	fmt.Printf("Me going to write some messages to kafka!")
	topic := "event.status"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	now := time.Now()
	message := fmt.Sprintf("%d-%d-%d-%d_%d_%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
		//kafka.Message{Value: []byte("two!")},
		//kafka.Message{Value: []byte("three!")},
	)

	fmt.Printf("Sent message: %s to kafka \n", message)
	conn.Close()
}

func main() {

	writeSome()
}
