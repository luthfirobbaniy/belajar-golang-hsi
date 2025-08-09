package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	writer := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "student.registered",
	}

	defer writer.Close()

	eventByte, eventByteErr := json.Marshal(Event{
		StudentId: time.Now().Format(time.TimeOnly),
		Name:      "Luthfi",
		Status:    "student.registered",
	})

	if eventByteErr != nil {
		log.Fatalf("Failed to read message: %v", eventByte)
	}

	log.Printf("Event Processed: %s", string(eventByte))

	err := writer.WriteMessages(
		context.TODO(),
		kafka.Message{
			Key:   []byte("Key-1"),
			Value: eventByte,
		},
	)

	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "student.registration_failed",
		GroupID: "student-service-group-1",
	})

	ctx, cancel := context.WithTimeout(
		context.Background(),
		60*time.Second, // It takes time to wait for the Consumer Group status to become "Stable"
	)

	defer cancel()

	msg, err := reader.ReadMessage(ctx)

	if err != nil {
		log.Fatalf("Payment success!")
	}

	log.Printf("Receive an event: %s", string(msg.Value))
}
