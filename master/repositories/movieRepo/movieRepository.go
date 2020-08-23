package movieRepo

import "netflixApp/master/models"

type MovieRepository interface {
	GetAllMovie() ([]*models.Movie, error)
	GetMovieById(id string) (*models.Movie, error)
	AddMovie(movie models.Movie) error
	UpdateMovie(id int, movie models.Movie) error
	DeleteMovie(id int) error
}
