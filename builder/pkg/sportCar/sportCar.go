// Package sportCar implement method to create car
package sportCar

import "fmt"

//Car part interface
type carPart interface {
	Append(string)
	Return() string
}

// Interface to work with car
type SportCar interface {
	StartNewCar() error
	BuildCarBody() error
	ReturnCar() (string, error)
}

type car struct {
	carPart carPart
}

// StartNewCar require for a car info and appends it to object
func (c *car) StartNewCar() (err error) {
	var insert string
	fmt.Print("Enter your car name: ")
	if _, err = fmt.Scan(&insert); err == nil {
		c.carPart.Append("Congratulation! You choose: " + insert + "\n")
	}
	return
}

// BuildCarBody will ask for more info and append it to the object
func (c *car) BuildCarBody() (err error) {
	var insert string
	fmt.Print("Enter body type: ")
	if _, err = fmt.Scan(&insert); err == nil {
		c.carPart.Append("You choose " + insert +" body type!\n")
		fmt.Print("Choose color: ")
		if _, err := fmt.Scan(&insert); err == nil {
			c.carPart.Append(insert + " is great color!\n")
			fmt.Print("Choose where you will ride it: ")
			if _, err := fmt.Scan(&insert); err == nil {
				c.carPart.Append(insert + " is great place to ride your new car!\n")
			}
		}

	}
	return
}

// ReturnCar returns a car
func (c *car) ReturnCar() (res string, err error) {
	if res = c.carPart.Return(); len(res) == 0 {
		err = fmt.Errorf(" Error! Can not create car")
		return
	}
	err = nil
	return
}

// NewSportCar instance
func NewSportCar(c carPart) SportCar {
	return &car{
		carPart: c,
	}
}
