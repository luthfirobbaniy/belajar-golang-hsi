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
		GroupID: "finance-service-group-1",
	})

	defer reader.Close()

	log.Println("Finance Service started..")

	for {
		msg, err := reader.ReadMessage(context.Background()) // Check about context later

		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		log.Printf("Receive an event: %s", string(msg.Value))

		event := &Event{}

		json.Unmarshal(msg.Value, event)

		isSuccess := rand.Intn(2) // Return random integer number, 0 or 1

		log.Println("Before", isSuccess)

		if isSuccess == 0 {
			event.Status = "student.registration_failed"
		} else {
			event.Status = "student.registration_validated"
		}

		writer := kafka.Writer{
			Addr:  kafka.TCP("localhost:9092"),
			Topic: event.Status,
		}

		eventByte, err := json.Marshal(event)

		if err != nil {
			log.Fatalf("Failed to encode event to JSON: %v", err)
		}

		log.Printf("Event processed: %s", string(eventByte))

		writer.WriteMessages(
			context.TODO(),
			kafka.Message{
				Key:   []byte("Key-1"),
				Value: eventByte,
			},
		)

	}
}
