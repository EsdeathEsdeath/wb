package main

import (
	"facade/pkg/actors"
	"facade/pkg/facade"
	"facade/pkg/producer"
	"facade/pkg/screenwriter"
	"facade/pkg/stuntman"
)

func main() {
	actors := actors.NewActors()
	producer := producer.NewProducer()
	screenwriter := screenwriter.NewScreenwriter()
	stuntman := stuntman.NewStuntman()
	filmSquad := facade.FilmSquad(screenwriter, producer, actors, stuntman)
	filmSquad.MakeGreatFilm()
}
