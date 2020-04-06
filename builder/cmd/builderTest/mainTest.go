package mainTest

import (
	"builder/pkg/product"
	"testing"

	"github.com/EsdeathEsdeath/wb_pr/tree/builder/builder/pkg/builder"
	"github.com/EsdeathEsdeath/wb_pr/tree/builder/builder/pkg/producy"
	"github.com/EsdeathEsdeath/wb_pr/tree/builder/builder/pkg/sportCar"
)

var testPass = `Congratulation! You choose: Supra
You choose Coupe body type
	Black is great color!
City is great place to ride your new car!
`

var testFail = ` Error! Can not create car`

func TestPassed(t *testing.T) {
	car := product.NewCar("")
	carBuilder := sportCar.NewSportCar(car)
	director := builder.NewDirector(carBuilder)
	out, err := director.BuildCar()
	if out != testPass && err == nil {
		t.Errorf("Test failed!\nExpected:\n%v", testPass)
	}
}

func TestError(t *testing.T) {
	car := product.NewCar("")
	carBuilder := sportCar.NewSportCar(car)
	_, err := carBuilder.ReturnCar()
	if err == nil {
		t.Errorf("Incorrect error handling!\nExpected:%s", testFail)
	}
}

