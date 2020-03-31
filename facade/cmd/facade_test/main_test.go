package facade_test

import (
	"testing"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/actors"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/facade"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/producer"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/screenwriter"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/stuntman"
)

var testEmpty = ` play at great film`
var testOne = `Stan Smith play at great film`

func TestOk(t *testing.T)	{
	actors := actors.NewActors()
	producer := producer.NewProducer()
	screenwriter := screenwriter.NewScreenwriter()
	stuntman := stuntman.NewStuntman()
	filmSquad := facade.FilmSquad(screenwriter, producer,actors, stuntman)
	// testing for empty string
	out := filmSquad.CreateCast("")
	if out != testEmpty {
		t.Errorf("Test failed!\nExpected:\n%v\nGot %v", testEmpty, out)
	}
	out = filmSquad.CreateCast("Stan Smith")
	if out != testOne	{
		t.Errorf("Test failed!\nExpected:\n%v\nGot:\n%v", testOne, out)
	}
}
