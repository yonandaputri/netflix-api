package tools

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetPathVar(key string, r *http.Request) string {
	param := mux.Vars(r)
	return param[key]
}
