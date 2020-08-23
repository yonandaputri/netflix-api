package main

import (
	"netflixApp/config"
	"netflixApp/master"
)

func main()  {
	db,_ := config.ConnectionDB()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}