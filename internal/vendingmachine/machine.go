package vendingmachine

type Coin interface {
	Mass() float32
	Diameter() float32
	Thickness() float32
}

// VendingMachine .
type VendingMachine struct {
}

// NewVendingMachine .
func NewVendingMachine() *VendingMachine {
	return &VendingMachine{}
}

// AcceptCoin .
func (v *VendingMachine) AcceptCoin(coin Coin) {

}

// SelectProduct .
func (v *VendingMachine) SelectProduct(product string) {

}

// MakeChange .
func (v *VendingMachine) MakeChange() {

}

// ReturnCoins .
func (v *VendingMachine) ReturnCoins() {

}

// Display .
func (v *VendingMachine) Display() string {
	return "ERROR"
}

// RestockProducts .
func (v *VendingMachine) RestockProducts() {
}
