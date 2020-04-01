//Package employee implement function to work with developer
package developer

type Developer struct {
	Income int
}

func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}


