package main_control

import (
	"log"
	"minidelivery/internal/adapter/delivery"
	"minidelivery/internal/contract"
)

// MainControl sends delivery`s to their duties
type MainControl struct {
	freeDeliveries  []contract.Deliver
	usedDeliveries  []contract.Deliver
	logger          *log.Logger
	freedDeliveries chan contract.Deliver
}

func New(count int, logger *log.Logger) *MainControl {
	var deliveries []contract.Deliver

	for i := 0; i < count; i++ {
		del := delivery.New(i + 1)
		deliveries = append(deliveries, del)
	}
	ch := make(chan contract.Deliver)
	return &MainControl{
		freeDeliveries:  deliveries,
		logger:          logger,
		freedDeliveries: ch,
	}
}
