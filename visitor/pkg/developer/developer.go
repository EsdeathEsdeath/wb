// Package employee implement function to work with developer
package developer

type bonusIncome interface {
	BonusForDev(income int) (res int)
}

type Developer interface {
	Accept() int
}

type developer struct {
	bonusIncome
	income int

}

// Accept and return calculation bonus income for developer
func (d *developer) Accept() (res int){
	res = d.bonusIncome.BonusForDev(d.income)
	return
}

// New developer instance
func NewDev(bonusIncome bonusIncome, income int) Developer {
	return &developer{
		bonusIncome:   bonusIncome,
		income: income,
	}
}
