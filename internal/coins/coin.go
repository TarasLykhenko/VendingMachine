package coins

type Coin struct {
	mass      float32
	diameter  float32
	thickness float32
}

// Creates a new coin
func NewCoin(mass float32, diameter float32, thickness float32) *Coin {
	return &Coin{
		mass:      mass,
		diameter:  diameter,
		thickness: thickness,
	}
}

// Mass of the coin in grams
func (c *Coin) Mass() float32 {
	return c.mass
}

// Diameter of the coin in mm
func (c *Coin) Diameter() float32 {
	return c.diameter
}

// Thickness of the coin in mm
func (c *Coin) Thickness() float32 {
	return c.thickness
}

