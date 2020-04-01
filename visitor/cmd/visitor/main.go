package main

import (
	"visitor/pkg/developer"
	"visitor/pkg/engineer"
	"visitor/pkg/visitor"
	"visitor/pkg/employee"
)

func main() {
	backend := developer.Developer{Income: 1000}
	admin := engineer.Engineer{Income: 2000}
	backend.Accept(calculIncome.CalculIncome{BonusRate: 20})
	admin.Accept(calculIncome.CalculIncome{BonusRate:10})

}
