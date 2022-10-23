package contract

import (
	"minidelivery/internal/entity"
)

type MainControl interface {
	Deliver(dest entity.Location) int
	HandleFreeDeliveries()
}

type Deliver interface {
	GetLocation() entity.Location
	GetNumber() int

	Move(dest entity.Location, freedDeliveries chan<- Deliver)
}
