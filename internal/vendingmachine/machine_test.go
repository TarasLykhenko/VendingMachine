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
