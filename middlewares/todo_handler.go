package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateTodoHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		description := c.FormValue("description")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		_, err := db.Exec("INSERT INTO todos (title, description, user_id) VALUES (?, ?, ?)", title, description, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal membuat todo"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Todo berhasil dibuat"})
	}
}
