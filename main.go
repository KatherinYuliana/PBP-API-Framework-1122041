package main

import (
	"Framework_API/controllers"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db := controllers.Connect()
    defer db.Close()
	
	e := echo.New()

	e.POST("/users", controllers.InsertUsers(db))
	e.GET("/users", controllers.GetAllUsers(db))
	e.PUT("/users/:id", controllers.UpdateUsers(db))
	e.DELETE("/users/:id", controllers.DeleteUsers(db))

	e.Start(":8888")
}