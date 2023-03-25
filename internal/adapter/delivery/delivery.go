package delivery

import "github.com/sajjad1993/minidelivery/internal/entity"

// Delivery moves to destination
type Delivery struct {
	Number   int
	location entity.Location
	free     bool
}

// IsFree free if  this delivery is  free returns true
func (d *Delivery) IsFree() bool {
	return d.free
}
func New(number int) *Delivery {
	return &Delivery{Number: number, free: true}
}
