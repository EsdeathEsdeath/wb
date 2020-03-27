package main

import (
	"facade/pkg/actors"
	"facade/pkg/facade"
	"facade/pkg/producer"
	"facade/pkg/screenwriter"
	"facade/pkg/stuntman"
)

func main() {
	a := actors.NewActors()
	p := producer.NewProducer()
	s := screenwriter.NewScreenwriter()
	st := stuntman.NewStuntman()
	f := facade.FilmSquad(s, p, a, st)
	f.MakeGreatFilm()
}
