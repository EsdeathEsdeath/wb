package main

import (
	"fmt"
	"os"

	"github.com/EsdeathEsdeath/wb_pr/tree/builder/builder/pkg/builder"
	"github.com/EsdeathEsdeath/wb_pr/tree/builder/builder/pkg/producy"
	"github.com/EsdeathEsdeath/wb_pr/tree/builder/builder/pkg/sportCar"
)
func main() {
	garage := "garage"
	g, err := os.Create(garage)
	if err == nil {
		car := product.NewCar("")
		var carBuilder sportCar.SportCar
		carBuilder = sportCar.NewSportCar(car)
		director := builder.NewDirector(carBuilder)
		if out, err := director.BuildCar(); err == nil {
			if _, err = g.WriteString(out); err != nil {
				fmt.Printf("Your car put in: %s", garage)
			}

		}
	}
}
