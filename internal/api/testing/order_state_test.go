package testing

import (
	"banking-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestOrderStateCancelled(t *testing.T) {
//	assert := assert.New(t)
//
//	var order, _ = entity.NewOrder(entity.Order{})
//
//	assert.Equal(entity.OrderAvailable, order.State)
//	assert.Len(order.StateHistory, 1)
//	assert.Equal(entity.OrderAvailable, order.StateHistory[0].State)
//
//	_, orderCancelledErr := order.SetNextState(entity.OrderCancelled)
//	assert.Nil(orderCancelledErr)
//	assert.Equal(entity.OrderCancelled, order.State)
//	assert.Len(order.StateHistory, 2)
//	assert.Equal(entity.OrderAvailable, order.StateHistory[0].State)
//	assert.Equal(entity.OrderCancelled, order.StateHistory[1].State)
//}

func TestOrderStateHappyWay(t *testing.T) {
	assert := assert.New(t)

	var order, _ = entity.NewOrder(entity.Order{})

	assert.Equal(entity.OrderAvailable, order.State)
	assert.Len(order.StateHistory, 1)
	assert.Equal(entity.OrderAvailable, order.StateHistory[0].State)

	order.State = entity.OrderAssigned

	_, orderPickedUpErr := order.SetNextState(entity.OrderTypeOrigin, entity.OrderReqSuccess)
	assert.Nil(orderPickedUpErr)
	assert.Equal(entity.OrderPickedUp, order.State)
	assert.Len(order.StateHistory, 2)
	assert.Equal(entity.OrderAvailable, order.StateHistory[0].State)
	assert.Equal(entity.OrderPickedUp, order.StateHistory[1].State)

	_, orderDeliveredErr := order.SetNextState(entity.OrderTypeDestination, entity.OrderReqSuccess)
	assert.Nil(orderDeliveredErr)
	assert.Equal(entity.OrderDelivered, order.State)
	assert.Len(order.StateHistory, 3)
	assert.Equal(entity.OrderAvailable, order.StateHistory[0].State)
	assert.Equal(entity.OrderPickedUp, order.StateHistory[1].State)
	assert.Equal(entity.OrderDelivered, order.StateHistory[2].State)
}
