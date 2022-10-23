package main_control

import (
	"context"
	"fmt"
	"minidelivery/internal/adapter/delivery"
	"minidelivery/internal/contract"
	"minidelivery/internal/entity"
)

// selectBestDelivery gets a location and returns its index
func (m *MainControl) selectBestDelivery(dest entity.Location) (int, error) {
	if len(m.freeDeliveries) == 0 {
		return -1, ErrorNoDelivery{}
	}
	for key, value := range m.freeDeliveries {
		tmp := value.GetLocation()
		if tmp.X == 0 && tmp.Y == 0 {
			return key, nil
		}
	}
	return m.getTheNearest(dest), nil
}

// getTheNearest returns the index of the nearest location
func (m *MainControl) getTheNearest(dest entity.Location) int {

	theNearest := 100000
	index := 0
	for i, value := range m.freeDeliveries {
		distance := delivery.CalculateDistance(value.GetLocation(), dest)
		if distance < theNearest {
			theNearest = distance
			index = i
		}
	}
	return index
}
func (m *MainControl) selectAndRemoveDelivery(dest entity.Location) contract.Deliver {
	// get index
	index, err := m.selectBestDelivery(dest)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//append to used deliveries
	selectedDelivery := m.freeDeliveries[index]
	m.usedDeliveries = append(m.usedDeliveries, selectedDelivery)
	// remove from free deliveries
	m.freeDeliveries = remove(m.freeDeliveries, index)
	fmt.Printf("free : %q \n used : %q  \n ", m.freeDeliveries, m.usedDeliveries)

	return selectedDelivery
}

// Deliver delivers the order
func (m *MainControl) Deliver(dest entity.Location, ctx context.Context) error {
	//select the best delivery
	deliveryGuy := m.selectAndRemoveDelivery(dest)
	if deliveryGuy != nil {
		go deliveryGuy.Move(dest, m.freedDeliveries)

	}
	return nil
}

// removes an item from slice //todo change it to generics
func remove(slice []contract.Deliver, s int) []contract.Deliver {
	return append(slice[:s], slice[s+1:]...)
}

func (m *MainControl) HandleFreeDeliveries() {
	fmt.Println("handle delivery now .... ")
	for value := range m.freedDeliveries {
		fmt.Printf("number:%d is free now \n ", value.GetNumber())
		for key, v := range m.usedDeliveries {
			if value.GetNumber() == v.GetNumber() {
				m.freeDeliveries = append(m.freeDeliveries, value)
				// remove from used deliveries
				m.usedDeliveries = remove(m.usedDeliveries, key)
				fmt.Printf("after delivery :  free : %q \n used : %q  \n ", m.freeDeliveries, m.usedDeliveries)

				break
			}
		}
	}
}
