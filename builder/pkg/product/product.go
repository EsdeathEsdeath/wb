package product

// Interface to work with complex object
type Car interface {
	Append(string)
	Return() string
}

type car struct {
	car     string
}

// Append appends to a site
func (c *car) Append(str string) {
	c.car += str
}

// Return returns a site
func (c *car) Return() string {
	return c.car
}

// NewSite ...
func NewCar(s string) Car {
	return &car{
		car:     "",
	}
}
