package vendingmachine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVendingMachine(t *testing.T) {
	expected := &VendingMachine{}

	vm := NewVendingMachine()

	assert.Equal(t, expected, vm)
}

func TestMakeChange(t *testing.T) {
	var expected float32
	expected = 1.35

	vm := NewVendingMachine()
	result := vm.MakeChange(2.00, 0.65)

	assert.Equal(t, expected, result)
}
