package main

import (
	"context"
	"encoding/json"
	"log"

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
		Topic:   "student.registration_validated",
		GroupID: "academic-service-group-1",
	})

	defer reader.Close()

	log.Println("Academic Service started..")

	for {
		msg, err := reader.ReadMessage(context.Background()) // Check about context later

		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		log.Printf("Receive an event: %s", string(msg.Value))

		event := &Event{}

		json.Unmarshal(msg.Value, event)

		event.Status = "student.academic_initialized"

		eventByte, err := json.Marshal(event)

		if err != nil {
			log.Fatalf("Failed to encode event to JSON: %v", err)
		}

		log.Printf("Event processed: %s", string(eventByte))
	}
}
