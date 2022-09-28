package entity_test

import (
	"testing"

	"github.com/eduardouliana/fullcycle-practice-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}

	assert.Error(t, order.IsValid(), "invalid ID")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123"}

	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 10}

	assert.Error(t, order.IsValid(), "invalid Tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShould_ReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)

	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)

	err = order.CalculateFinalPrice()
	assert.NoError(t, err)

	assert.Equal(t, 12.0, order.FinalPrice)

}
