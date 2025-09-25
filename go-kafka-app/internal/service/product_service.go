package service

import "go-kafka-app/internal/app"

func ProcessProductEvent(event app.ProductEvent){
	println("Processd:", event.ID, event.Name, event.Price)
}