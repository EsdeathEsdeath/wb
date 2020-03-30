package main

import (
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/actors"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/facade"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/producer"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/screenwriter"
	"github.com/EsdeathEsdeath/wb_pr/facade/pkg/stuntman"
)

func main() {
	actors := actors.NewActors()
	producer := producer.NewProducer()
	screenwriter := screenwriter.NewScreenwriter()
	stuntman := stuntman.NewStuntman()
	filmSquad := facade.FilmSquad(screenwriter, producer, actors, stuntman)
	filmSquad.MakeGreatFilm()
}
