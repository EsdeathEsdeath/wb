//Package employee implement function to work with engineer
package engineer

type Engineer struct {
	Income int
}

func (e Engineer) Accept(v Visitor) {
	v.VisitEngineer(e)
}
