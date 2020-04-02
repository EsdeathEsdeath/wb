// Package engineer implement function to work with engineer
package engineer

type bonusIncome interface {
	BonusForEng(income int) (res int)
}

type Engineer interface {
	Accept() int
}

type engineer struct {
	bonusIncome
	income int

}

// Accept and return calculation bonus income for engineer
func (e *engineer) Accept() (res int){
	res = e.bonusIncome.BonusForEng(e.income)
	return
}

// New dev instance
func NewEng(bonusIncome bonusIncome, income int) Engineer {
	return &engineer{
		bonusIncome:   bonusIncome,
		income: income,
	}
}
