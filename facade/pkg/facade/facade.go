//Package facade implement method to create a movie
package facade

import "fmt"

type filmScreenwriter interface {
	WriteScript(name string)
}

type filmProducers interface {
	MakeMovie(name string)
}

type filmActors interface {
	SetCast(cast string) string
}

type filmStuntman interface {
	DoTricks(name string)
}

// Film Interface to work with film type
type Film interface {
	CreateCast(string) string
	MakeGreatFilm()
}

type film struct {
	screenwriter   filmScreenwriter
	producer   filmProducers
	actors filmActors
	stuntman  filmStuntman
}

//Function to create your own actors cast
func (f *film) CreateCast(cast string) string{
	res := f.actors.SetCast(cast)
	fmt.Println(res)
	return res
}

//Set group to create a great film
func (f *film) MakeGreatFilm(){
	f.screenwriter.WriteScript("Billy Wilder")
	f.producer.MakeMovie("Michael Mann")
	f.stuntman.DoTricks("Dar Robinson")
}

// Film squad creating
func FilmSquad(
	screenwriter filmScreenwriter,
	producer filmProducers,
	actors filmActors,
	stuntman filmStuntman) Film {
	return &film{
		screenwriter:   screenwriter,
		producer:   producer,
		actors: actors,
		stuntman:  stuntman,
	}
}

