package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReturnAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}
func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReturnAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}
func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReturnAnError(t *testing.T) {
	order := Order{ID: "123", Price: 3}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenShouldReceiveCreateOrderWithParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 8.0,
		Tax:   1.0,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 8.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	assert.Nil(t, order.IsValid())
}
func TestGivenAValidParams_WhenICallNewOrderFunc_ThenShouldReceiveCreateOrderWithParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)

}
func TestGivenANotValidParams_WhenICallNewOrderFunc_ThenShouldReceiveAnError(t *testing.T) {
	order, err := NewOrder("123", 10.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)

}

func TestGivenPriceTax_WhenICallCalculatePrice_ThanItShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.2, 1.8)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, order.FinalPrice, 12.0)

}
