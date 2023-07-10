package config

import (
	"database/sql"
	"example/blog/helper"
	"fmt"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbName   = "blog"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("mysql", sqlInfo)

	helper.PanicIfErrors(err)

	err = db.Ping()

	helper.PanicIfErrors(err)

	log.Info().Msg("Connected to database!!")

	return db
}
