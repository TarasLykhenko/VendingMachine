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
	t.Run("Make change input 2.00, product 0.55", func(t *testing.T) {
		v := VendingMachine{

			amoney: 2.00,
		}
		v.MakeChange(0.55)

		expected = 1.45

		assert.Equal(t, expected, v.amoney)
	})
}

func TestReturnCoins(t *testing.T) {
	t.Run("The machine has no coins", func(t *testing.T) {
		v := VendingMachine{
			amoney: 1.00,
			status: Status{
				quarter: 0,
				dime:    0,
				nickel:  0,
			},
		}
		status, err := v.ReturnCoins()

		assert.Nil(t, status)
		assert.Error(t, err)

	})
	t.Run("The machine has no money and no coins", func(t *testing.T) {
		v := VendingMachine{
			amoney: 0.00,
			status: Status{
				quarter: 0,
				dime:    0,
				nickel:  0,
			},
		}
		status, err := v.ReturnCoins()
		assert.Equal(t, Status{}, *status)
		assert.NotNil(t, status)
		assert.NoError(t, err)

	})
}

func TestRetriveMoney(t *testing.T) {
	t.Run("The machine has no coins to retrive", func(t *testing.T) {
		v := VendingMachine{
			status: Status{
				quarter: 3,
				dime:    3,
				nickel:  3,
			},
		}
		status, err := v.RetriveMoney()

		retriveCoins := Status{
			nickel:  0,
			dime:    0,
			quarter: 0,
		}
		assert.Equal(t, retriveCoins, *status)
		assert.Nil(t, err)
	})
	t.Run("The machine has coins to retrive", func(t *testing.T) {
		v := VendingMachine{
			status: Status{
				quarter: 6,
				dime:    6,
				nickel:  6,
			},
		}
		status, err := v.RetriveMoney()

		retriveCoins := Status{
			nickel:  2,
			dime:    2,
			quarter: 2,
		}
		assert.Equal(t, retriveCoins, *status)
		assert.Nil(t, err)
	})
}
