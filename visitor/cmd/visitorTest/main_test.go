// Package visitorTest implement test to check program output
package visitorTest

import (
	"fmt"
	"testing"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/calculIncome"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/developer"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/engineer"
	"github.com/EsdeathEsdeath/wb_pr/tree/visitor_first/visitor/pkg/visitor"
)

var devRes, enginRes = 7666, 22000

// Test developer value
func TestForBackOk(t *testing.T) {
	bonus := calculIncome.NewInc()
	backend := developer.NewDev(bonus, 1000)
	visit := visitor.NewVisitor(backend)
	bonusForBack := visit.Accept()

	if bonusForBack != devRes {
		fmt.Errorf("Error!\nexpected:\n%d\ngot:\n%d", devRes, bonusForBack)
	}
}

// Test engineer value
func TestForEngOk(t *testing.T) {
	bonus := calculIncome.NewInc()
	chief := engineer.NewEng(bonus, 2000)
	visit := visitor.NewVisitor(chief)
	bonusForBack := visit.Accept()

	if bonusForBack != enginRes {
		fmt.Errorf("Error!\nexpected:\n%d\ngot:\n%d", devRes, bonusForBack)
	}
}

