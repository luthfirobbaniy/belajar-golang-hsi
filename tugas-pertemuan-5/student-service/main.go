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
	// Publish to "student.registered"
	writer := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "student.registered",
	}

	defer writer.Close()

	event := &Event{
		StudentId: "1",
		Name:      "Luthfi",
		Status:    "student.registered",
	}

	eventByte, err := json.Marshal(event)

	if err != nil {
		log.Fatalf("[Student Service] Failed to read message: %v", eventByte)
	}

	writeMessageErr := writer.WriteMessages(
		context.TODO(),
		kafka.Message{
			Key:   []byte("Key-1"),
			Value: eventByte,
		},
	)

	if writeMessageErr != nil {
		log.Fatalf("[Student Service] Failed to write message: %v", writeMessageErr)
	}

	log.Printf("[Student Service] Sent event: %s", event.Status)

	// Subscribe to student.registration_failed
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

	_, readMessageErr := reader.ReadMessage(ctx)

	if readMessageErr != nil {
		log.Fatalf("[Student Service] Payment success!")
	}

	log.Printf("[Student Service] Received event: %s, student_id: %s", event.Status, event.StudentId)
}
