// Package producer implement method to filming
package producer

import "fmt"

//  Make movie
type Producer interface {
	MakeMovie(name string)
}

type producer struct {
	name string
}

// Filming
func (p *producer) MakeMovie(name string) {
	fmt.Println(name, "make a great movie")
}

// Create producer
func NewProducer() Producer {
	return &producer{}
}
