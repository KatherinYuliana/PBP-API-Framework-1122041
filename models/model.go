package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}