package config

import (
	"database/sql"
	"fmt"
	"log"
	"netflixApp/tools"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*sql.DB, error) {
	dbUser := tools.ReadEnv("DB_USER", "root")
	dbPass := tools.ReadEnv("DB_PASS", "Viontin@12")
	dbHost := tools.ReadEnv("DB_HOST", "localhost")
	dbPort := tools.ReadEnv("DB_PORT", "3306")
	dbName := tools.ReadEnv("DB_NAME", "schema")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}

	return db, nil
}
