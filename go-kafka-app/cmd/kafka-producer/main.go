package main

import (
	"context"
	"go-kafka-app/internal/app"
	"go-kafka-app/internal/config"
	"go-kafka-app/internal/kafka"
	"time"
)

func main() {
	cfg := config.Load()
	producer := kafka.NewProducer(cfg.Broker,cfg.Topic)
	defer producer.Close()

	event := app.ProductEvent{ID: "1001", Name: "Tata", Price: 9.7}
	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := producer.Produce(ctx,event); err != nil {
		panic(err)
	}
}