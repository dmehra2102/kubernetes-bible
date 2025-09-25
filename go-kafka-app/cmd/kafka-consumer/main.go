package main

import (
	"context"
	"go-kafka-app/internal/config"
	"go-kafka-app/internal/kafka"
	"go-kafka-app/internal/service"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	kafka.StartConsumer(ctx, cfg.Broker,cfg.Topic,cfg.Group, service.ProcessProductEvent)
}