package movieUsecase

import "netflixApp/master/models"

type MovieUseCase interface {
	GetMovies() ([]*models.Movie, error)
	GetMovieById(id string) (*models.Movie, error)
	PostMovie(movie models.Movie) error
	PutMovie(id int, movie models.Movie) error
	DeleteMovie(id int) error
}
