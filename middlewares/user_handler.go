package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		role := c.FormValue("role")

		_, err := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", username, password, role)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membuat user"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "User berhasil dibuat"})
	}
}
