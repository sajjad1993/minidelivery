package contract

import (
	"github.com/sajjad1993/minidelivery/internal/entity"
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
