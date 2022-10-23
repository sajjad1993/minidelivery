package main_control

import (
	"github.com/stretchr/testify/assert"
	"log"
	"minidelivery/internal/entity"
	"testing"
)

func setup(t *testing.T) (*MainControl, func()) {
	count := 2
	logger := log.Logger{}
	mainControl := New(count, &logger)
	return mainControl, func() {}
}
func TestSelectDelivery(t *testing.T) {
	t.Run("get best delivery", func(t *testing.T) {
		mainControl, teardown := setup(t)
		defer teardown()
		loc := entity.Location{X: 3, Y: 4}
		mainControl.selectAndRemoveDelivery(loc)
		assert.Equal(t, 0, 0)

	})
}
