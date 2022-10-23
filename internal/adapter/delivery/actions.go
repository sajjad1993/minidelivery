package delivery

import (
	"fmt"
	"minidelivery/internal/contract"
	"minidelivery/internal/entity"
	"time"
)

func (d *Delivery) GetLocation() entity.Location {
	return d.location
}
func (d *Delivery) GetNumber() int {
	return d.Number
}

// Move  gets a dest and goes  there
func (d *Delivery) Move(dest entity.Location, freedDeliveries chan<- contract.Deliver) {
	// set his self unfree
	d.free = false
	// start moving to destination
	distance := d.location.Distance(dest)

	for i := 0; i < distance; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("i`m number:%d and i`m in way \n ", d.Number)
	}
	d.location.X = dest.X
	d.location.Y = dest.Y
	d.free = true
	fmt.Printf("i`m number:%d and i`ve reccived and my location is ( %d , %d ) \n ", d.Number, d.location.X, d.location.Y)

	freedDeliveries <- d

}
