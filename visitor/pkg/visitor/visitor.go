// Package visitor implement interface to work with visitor
package visitor


type BonusIncome interface {
	Accept() int
}

type Visitor interface {
	Accept() int
}

type visitor struct {
	bonusIncome BonusIncome
}

// Accept returns bonus income for employee
func (v *visitor) Accept() (res int) {
	res = v.bonusIncome.Accept()
	return
}

// NewVisitor instance
func NewVisitor(b BonusIncome) Visitor {
	return &visitor{bonusIncome: b}
}
