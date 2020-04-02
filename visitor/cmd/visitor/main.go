package main

import (
	"fmt"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/calculIncome"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/developer"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/engineer"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/visitor"
)

func main() {
	bonus := calculIncome.NewInc()
	backend := developer.NewDev(bonus, 1000)
	visit := visitor.NewVisitor(backend)
	bonusForBack := visit.Accept()
	fmt.Println("New income for developer: ",bonusForBack)

	chief := engineer.NewEng(bonus, 2000)
	visit = visitor.NewVisitor(chief)
	bonusForChief := visit.Accept()
	fmt.Println("New income for engineer: ",bonusForChief)

}
