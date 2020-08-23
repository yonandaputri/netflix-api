package master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"netflixApp/master/controllers"
	"netflixApp/master/repositories/movieRepo"
	"netflixApp/master/usecases/movieUsecase"
)

func Init(r *mux.Router, db *sql.DB)  {
	movieRepo := movieRepo.InitMovieRepoImpl(db)
	movieUsecase := movieUsecase.InitMovieUseCase(movieRepo)
	controllers.MovieController(r, movieUsecase)
}