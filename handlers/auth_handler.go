package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/MIKTI_Tugas4_adi/middlewares"
	"github.com/MIKTI_Tugas4_adi/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		var user models.User
		err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username).
			Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil || user.Password != password {
			return echo.ErrUnauthorized
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["username"] = user.Username
		claims["role"] = user.Role
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString(middlewares.JwtSecret)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{"token": t})
	}
}
