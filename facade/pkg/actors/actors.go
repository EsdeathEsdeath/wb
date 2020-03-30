//Package actors implement method to play a role in movie
package actors

import "fmt"

// Play role
type Actors interface {
	PlayRole(name string)
}

type actors struct {
	name string
}

// Played their roles in the film
func (a *actors) PlayRole(name string) {
	fmt.Println(name, "play a great role")
}

// Create new actors
func NewActors() Actors {
	return &actors{}
}
