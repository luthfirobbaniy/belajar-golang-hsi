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

	log.Println("[Academic Service] Started..")

	for {
		msg, err := reader.ReadMessage(context.Background()) // Check about context later

		if err != nil {
			log.Fatalf("[Academic Service] Failed to read message: %v", err)
		}

		event := &Event{}

		json.Unmarshal(msg.Value, event)

		log.Printf("[Academic Service] Received event: %s, student_id: %s", event.Status, event.StudentId)

		event.Status = "student.academic_initialized"

		log.Printf("[Academic Service] Academic initialized for student_id: %s", event.StudentId)
	}
}
