// Package employee implement interface to work with employees
package employee

type Employee interface {
	Accept(Visitor)
}
