// Package calculIncome implement function to calculate income for employees
package calculIncome

type BonusIncome interface {
	BonusForDev(income int) (res int)
	BonusForEng(income int) (res int)
}

type calculIncome struct {}

// Function for calculate income for developer
func (c *calculIncome) BonusForDev(income int) (res int) {
	res = income + income * income/150
	return
}

// Function for calculate income for engineer
func (c *calculIncome) BonusForEng(income int) (res int) {
	res = income + income * income/200
	return
}

// New bonus income instance
func NewInc() BonusIncome {
	return &calculIncome{}
}
