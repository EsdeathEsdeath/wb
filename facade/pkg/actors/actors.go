//Package actors implement method to play a role in movie
package actors

// Play role
type Actors interface {
	SetCast (string) string
}

type actors struct {
	name []string
}

// SetCast appends actors and returns cast of film
//You can choose cast by yourself
func (a *actors) SetCast(cast string) string {
	var iter int
	var castList, buffer string
	if len(cast) != 0 {
		a.name = append(a.name, cast)
	}
	for iter, buffer = range a.name {
		if iter == 0 {
			castList += "" + buffer
		} else {
			castList += ", " + buffer
		}
	}
	return castList + " play at great film"
}

// Create new actors
func NewActors() Actors {
	return &actors{}
}
