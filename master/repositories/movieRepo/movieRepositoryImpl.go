package movieRepo

import (
	"database/sql"
	"fmt"
	"netflixApp/master/models"
	"netflixApp/utils"
)

type MovieRepoImpl struct {
	db *sql.DB
}

func (m MovieRepoImpl) GetAllMovie() ([]*models.Movie, error) {
	var movies []*models.Movie
	data, err := m.db.Query(utils.SELECT_MOVIE)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var movie = new(models.Movie)
		var err = data.Scan(&movie.Id, &movie.Title, &movie.Duration, &movie.ImageUrl, &movie.Synopsis)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (m MovieRepoImpl) GetMovieById(id string) (*models.Movie, error) {
	var movie = new(models.Movie)
	err := m.db.QueryRow(utils.SELECT_MOVIE_BY_ID, id).Scan(&movie.Id, &movie.Title, &movie.Duration, &movie.ImageUrl, &movie.Synopsis)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m MovieRepoImpl) AddMovie(movie models.Movie) error {
	data, err := m.db.Begin()

	if err != nil {
		return err
	}

	row, err := m.db.Prepare(utils.INSERT_MOVIE)

	if err != nil {
		return err
	}

	_, err = row.Exec(&movie.Title, &movie.Duration, &movie.ImageUrl, &movie.Synopsis)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (m MovieRepoImpl) UpdateMovie(id int, movie models.Movie) error {
	data, err := m.db.Begin()
	if err != nil {
		return err
	}

	row, err := m.db.Prepare(utils.UPDATE_MOVIE)
	if err != nil {
		return err
	}

	_, err = row.Exec(&movie.Title, &movie.Duration, &movie.ImageUrl, &movie.Synopsis, id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (m MovieRepoImpl) DeleteMovie(id int) error {
	data, err := m.db.Begin()
	if err != nil {
		return err
	}

	row, err := m.db.Prepare(utils.DELETE_MOVIE)
	if err != nil {
		return err
	}

	_, err = row.Exec(id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func InitMovieRepoImpl(db *sql.DB) MovieRepository {
	return &MovieRepoImpl{db}
}
