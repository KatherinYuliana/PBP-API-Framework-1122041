package main

import (
	"Framework_API/controllers"
	//"database/sql"
	"fmt"
	"log"
	"net/http"

	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db := controllers.Connect()
	// db, err := controllers.Connect()
    // if err != nil {
    //     log.Fatal(err)
    // }
    defer db.Close()
	// defer db.Close()
	e := echo.New()

	e.POST("/users", controllers.InsertUsers(db))

	// e.POST("/users", controllers.InsertUsers())
	e.GET("/users", controllers.GetAllUsers(db))
	e.PUT("/users/:id", controllers.UpdateUsers(db))
	e.DELETE("/users/:id", controllers.DeleteUsers(db))

	// e.Logger.Fatal(e.Start(":8888"))
	e.Start(":8888")
	http.Handle("/", e)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
}