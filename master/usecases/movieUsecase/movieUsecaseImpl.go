package movieUsecase

import (
	"netflixApp/master/models"
	"netflixApp/master/repositories/movieRepo"
)

type MovieUsecaseImpl struct {
	movieRepo movieRepo.MovieRepository
}

func (m MovieUsecaseImpl) GetMovies() ([]*models.Movie, error) {
	movies, err := m.movieRepo.GetAllMovie()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (m MovieUsecaseImpl) GetMovieById(id string) (*models.Movie, error) {
	movie, err := m.movieRepo.GetMovieById(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m MovieUsecaseImpl) PostMovie(movie models.Movie) error {
	err := m.movieRepo.AddMovie(movie)
	if err != nil {
		return err
	}
	return nil
}

func (m MovieUsecaseImpl) PutMovie(id int, movie models.Movie) error {
	err := m.movieRepo.UpdateMovie(id, movie)
	if err != nil {
		return err
	}
	return nil
}

func (m MovieUsecaseImpl) DeleteMovie(id int) error {
	err := m.movieRepo.DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}

func InitMovieUseCase(movieRepo movieRepo.MovieRepository) MovieUseCase {
	return &MovieUsecaseImpl{movieRepo}
}
