package contract

import (
	"context"
	"minidelivery/internal/entity"
)

type MainControl interface {
	Deliver(dest entity.Location, ctx context.Context) error
	HandleFreeDeliveries()
}

type Deliver interface {
	GetLocation() entity.Location
	GetNumber() int

	Move(dest entity.Location, freedDeliveries chan<- Deliver)
}
