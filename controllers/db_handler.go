package controllers

import (
	"database/sql"
	//_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

func Connect() *sql.DB{
	e := echo.New()
	//db, err := sql.Open("sqlite3", "lat_pbp_2")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/lat_pbp_2?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		// log.Fatal(err)
		//return nil, err
		e.Logger.Fatal(err)
	}
	return db
	//defer db.Close()
}