// Package screenwriter implement method to  write a script
package screenwriter

import "fmt"

// Script writing
type Screenwriter interface {
	WriteScript(name string)
}

type screenwriter struct {
	name string
}

// The script is written
func (s *screenwriter) WriteScript(name string) {
	fmt.Println(name, "write a great script")
}

// Creating new screenwriter
func NewScreenwriter() Screenwriter {
	return &screenwriter{}
}
