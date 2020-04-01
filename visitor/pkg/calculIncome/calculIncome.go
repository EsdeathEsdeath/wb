//
package calculIncome

import (
"fmt"
)

type CalculIncome struct {
	BonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.BonusRate/100)
}
func (c CalculIncome) VisitEngineer(e Enginner) {
	fmt.Println(e.Income + e.Income*c.BonusRate/100)
}
