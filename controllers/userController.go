package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	m "Framework_API/models"
)

// Insert User
func InsertUser(db *sql.DB, name string, age int) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func InsertUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(m.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		_, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, user)
	}
}

// Get All User
func GetAllUser(db *sql.DB) ([]m.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []m.User
	for rows.Next() {
		var user m.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetAllUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := GetAllUser(db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}
}

// Update User
func UpdateUser(db *sql.DB, id int, name string, age int) (int64, error) {
	result, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", name, age, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func UpdateUsers (db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user := new(m.User)
		if err := c.Bind(user); err != nil {
			return err
		}
		_, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", user.Name, user.Age, id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}


// Delete User
func DeleteUser(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func DeleteUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := db.Exec("DELETE FROM users WHERE id=?", id)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	}
}