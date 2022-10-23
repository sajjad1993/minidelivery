package main_control

import (
	"github.com/stretchr/testify/assert"
	"log"
	"minidelivery/internal/entity"
	"testing"
	"time"
)

func setup(t *testing.T) (*MainControl, func()) {
	count := 2
	logger := log.Logger{}
	mainControl := New(count, &logger)
	return mainControl, func() {}
}
func TestSelectDelivery(t *testing.T) {
	t.Run("send 2 delivery and must delivery number 2 get the next job !!", func(t *testing.T) {
		mainControl, teardown := setup(t)
		go mainControl.HandleFreeDeliveries()
		defer teardown()
		loc1 := entity.Location{X: 1, Y: 1}
		loc2 := entity.Location{X: 2, Y: 2}
		loc3 := entity.Location{X: 3, Y: 3}
		deliveryNumber := mainControl.Deliver(loc1)
		assert.NotEqual(t, deliveryNumber, -1, "we must`ve already had a free delivery guy")
		deliveryNumber = mainControl.Deliver(loc2)
		assert.NotEqual(t, deliveryNumber, -1, "we must`ve already had a free delivery guy")
		time.Sleep(5 * time.Second)
		deliveryNumber = mainControl.Deliver(loc3)

		assert.Equal(t, deliveryNumber, 2, "number 2 is near than number 1 .. something went wrong ")

	})
}
