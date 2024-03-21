package controllers

import (
	"database/sql"
	

	"github.com/labstack/echo/v4"
)

func Connect() *sql.DB{
	e := echo.New()
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/lat_pbp_2?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		e.Logger.Fatal(err)
	}
	return db
}