package kafka

import (
	"context"
	"encoding/json"
	"go-kafka-app/internal/app"

	"github.com/segmentio/kafka-go"
)

type EventHandlerFunc func(app.ProductEvent)

func StartConsumer(ctx context.Context, broker,topic, group string, handler EventHandlerFunc)error{
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic: topic,
		GroupID: group,
	})
	defer r.Close()

	for {
		m,err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		var event app.ProductEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			handler(event)
		}
	}
	return nil
}