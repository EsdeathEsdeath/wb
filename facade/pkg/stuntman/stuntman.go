package stuntman

import "fmt"

// Performs tricks for the film
type Stuntman interface {
	DoTricks(name string)
}

type stuntman struct {
	name string
}

// Do some tricks
func (s *stuntman) DoTricks(name string) {
	fmt.Println(name, "do a great tricks")
}

// New stuntman creating
func NewStuntman() Stuntman {
	return &stuntman{}
}
