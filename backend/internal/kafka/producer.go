package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	topic  string
}

func NewProducer(brokers []string, topic string) *Producer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10,
	}

	return &Producer{
		writer: w,
		topic:  topic,
	}
}

func (p *Producer) Publish(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Published message to topic %s with key %s", p.topic, key)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

type Consumer struct {
	reader *kafka.Reader
	topic  string
	group  string
}

func NewConsumer(brokers []string, topic string, groupID string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &Consumer{
		reader: r,
		topic:  topic,
		group:  groupID,
	}
}

func (c *Consumer) Consume(ctx context.Context, handler func(message kafka.Message) error) error {
	log.Printf("Starting consumer for topic %s with group %s", c.topic, c.group)

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			log.Printf("Error reading message: %v", err)
			continue
		}

		if err := handler(msg); err != nil {
			log.Printf("Error handling message: %v", err)
			continue
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
