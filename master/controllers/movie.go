package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"netflixApp/master/models"
	"netflixApp/master/usecases/movieUsecase"
	"netflixApp/tools"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	MovieUseCase movieUsecase.MovieUseCase
}

func MovieController(r *mux.Router, service movieUsecase.MovieUseCase) {
	movieHandler := MovieHandler{service}
	r.HandleFunc("/movies", movieHandler.ListMovies).Methods(http.MethodGet)
	r.HandleFunc("/movie/{id}", movieHandler.GetMovieById).Methods(http.MethodGet)
	r.HandleFunc("/movie", movieHandler.PostMovie).Methods(http.MethodPost)
	r.HandleFunc("/movie/{id}", movieHandler.PutMovie).Methods(http.MethodPut)
	r.HandleFunc("/movie/{id}", movieHandler.DeleteMovie).Methods(http.MethodDelete)
}

func (mh MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.MovieUseCase.GetMovies()

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfMovies, err := json.Marshal(movies)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfMovies))
	fmt.Println("Endpoint hit: Get Movies")
}

func (mh MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := tools.GetPathVar(key, r)
	movie, err := mh.MovieUseCase.GetMovieById(id)

	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfMovie, err := json.Marshal(movie)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfMovie))
	fmt.Println("Endpoint hit: Get Movie By Id")
}

func (mh MovieHandler) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	err = mh.MovieUseCase.PostMovie(movie)
	if err != nil {
		w.Write([]byte("Cannot add data"))
	} else {
		data, _ := json.Marshal(movie)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(data))
		fmt.Println("Endpoint hit: Post Movie")
	}
}

func (mh MovieHandler) PutMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	key := "id"
	id := tools.GetPathVar(key, r)
	idMovie, _ := strconv.Atoi(id)
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	} else {
		err = mh.MovieUseCase.PutMovie(idMovie, movie)
		if err != nil {
			w.Write([]byte("Cannot update data"))
		} else {
			data, _ := json.Marshal(movie)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(data))
			fmt.Println("Endpoint hit: Put Room")
		}
	}
}

func (mh MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	key := "id"
	id := tools.GetPathVar(key, r)
	idMovie, _ := strconv.Atoi(id)
	err := mh.MovieUseCase.DeleteMovie(idMovie)
	if err != nil {
		w.Write([]byte("Cannot delete data"))
	} else {
		w.Write([]byte("Succes delete data"))
		fmt.Println("Endpoint hit: Delete Room")
	}
}
