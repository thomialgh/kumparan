package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// InitWriter -
func InitWriter(address string, topic string, clientID string) kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers: []string{address},
		Topic:   topic,
		Dialer: &kafka.Dialer{
			ClientID:        clientID,
			TransactionalID: "test",
		},
		BatchTimeout: 1 * time.Millisecond,
		BatchSize:    1,
		ReadTimeout:  300 * time.Millisecond,
		WriteTimeout: 300 * time.Millisecond,
		Balancer:     &kafka.Hash{},
	}
}

// Publish -
func Publish(config kafka.WriterConfig, key []byte, msg interface{}) error {
	writer := kafka.NewWriter(config)
	defer writer.Close()
	fmt.Println("test")
	msgByte, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: msgByte,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// InitReader -
func InitReader(address, topic, consumerGroup string) kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers: []string{address},
		Topic:   topic,
		GroupID: consumerGroup,
		MaxWait: 200 * time.Millisecond,
	}
}

// Cousume -
func Cousume(config kafka.ReaderConfig) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reader := kafka.NewReader(config)
	defer reader.Close()

	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		return []byte{}, err
	}

	return msg.Value, err

}
