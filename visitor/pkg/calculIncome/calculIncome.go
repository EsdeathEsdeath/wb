// Package calculIncome implement function to calculate income for both visit
package calculIncome

import (
"fmt"
)

type CalculIncome struct {
	BonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income * c.BonusRate/10)
}
func (c CalculIncome) VisitEngineer(e Enginner) {
	fmt.Println(e.Income + e.Income * c.BonusRate/100)
}
