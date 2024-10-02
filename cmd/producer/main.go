package main

import (
	"context"
	"encoding/json"
	"github.com/prok05/wb-level-0/config"
	"github.com/prok05/wb-level-0/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", config.Envs.KafkaURL, config.Envs.KafkaTopic, 0)
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalf("failed to set write deadline: %v", err)
	}

	order := utils.GenerateOrder()
	marshalledOrder, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("failed to marshal order: %v", err)
	}

	_, err = conn.Write(marshalledOrder)
	if err != nil {
		log.Fatalf("failed to write message: %v", err)
	}
	log.Println("Message was added successfully")

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
