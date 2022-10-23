package delivery

import (
	"github.com/stretchr/testify/assert"
	"minidelivery/internal/contract"
	"minidelivery/internal/entity"
	"testing"
	"time"
)

func setup(t *testing.T) (*Delivery, func()) {
	number := 1
	deliveryGuy := New(number)
	return deliveryGuy, func() {}
}
func TestMoveDelivery(t *testing.T) {
	t.Run("send 2 delivery and must delivery number 2 get the next job !!", func(t *testing.T) {
		deliveryGuy, teardown := setup(t)
		defer teardown()
		loc1 := entity.Location{X: 1, Y: 1}
		ch := make(chan contract.Deliver)
		go deliveryGuy.Move(loc1, ch)
		time.Sleep(2 * time.Second)
		for item := range ch {
			assert.Equal(t, item.GetNumber(), 1, "number must be 1 .. something went wrong ")
			close(ch)
		}

	})
}
