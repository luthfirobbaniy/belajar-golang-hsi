package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "student.registered",
		GroupID: "finance-service-group",
	})

	defer reader.Close()

	log.Println("[Finance Service] Started..")

	for {
		msg, err := reader.ReadMessage(context.Background()) // Check about context later

		if err != nil {
			log.Fatalf("[Finance Service] Failed to read message: %v", err)
		}

		event := &Event{}

		json.Unmarshal(msg.Value, event)

		log.Printf("[Finance Service] Received event: %s, student_id: %s", event.Status, event.StudentId)

		isSuccess := rand.Intn(2) // Return random integer number, 0 or 1

		if isSuccess == 0 {
			event.Status = "student.registration_failed"
		} else {
			event.Status = "student.registration_validated"
			log.Printf("[Finance Service] Payment validated for student_id: %s", event.StudentId)
		}

		writer := kafka.Writer{
			Addr:  kafka.TCP("localhost:9092"),
			Topic: event.Status,
		}

		eventByte, err := json.Marshal(event)

		if err != nil {
			log.Fatalf("[Finance Service] Failed to encode event to JSON: %v", err)
		}

		log.Printf("[Finance Service] Sent event: %s", event.Status)

		writeMessageErr := writer.WriteMessages(
			context.TODO(),
			kafka.Message{
				Key:   []byte("Key-1"),
				Value: eventByte,
			},
		)

		if writeMessageErr != nil {
			log.Fatalf("[Finance Service] Failed to write message: %v", writeMessageErr)
		}
	}
}
