// Package builder implement method to work with builder
package builder

// carBuilder interface to work with builder
type carBuilder interface {
	StartNewCar() error
	BuildCarBody() error
	ReturnCar() (string, error)
}

// Director interface to work with builder
type Director interface {
	BuildCar() (string, error)
}

type director struct {
	builder carBuilder
}

// BuildCar calls necessary func to do a car
func (d *director) BuildCar() (res string, err error) {
	if err = d.builder.StartNewCar(); err == nil {
		if err = d.builder.BuildCarBody(); err == nil {
			res, err = d.builder.ReturnCar()
		}
	}
	return
}

// NewDirector instance
func NewDirector(s carBuilder) Director {
	return &director{
		builder: s,
	}
}
