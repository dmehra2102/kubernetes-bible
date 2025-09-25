package kafka

import (
	"context"
	"encoding/json"
	"go-kafka-app/internal/app"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker, topic string) *Producer {
	w := &kafka.Writer{
		Addr: kafka.TCP(broker),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{writer: w}
}

func (p *Producer) Produce(ctx context.Context, event app.ProductEvent) error {
	val,err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{Key: []byte(event.ID),Value: val}
	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}