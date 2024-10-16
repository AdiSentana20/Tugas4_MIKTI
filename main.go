package main

import (
	"github.com/MIKTI_Tugas4_adi/config"
	"github.com/MIKTI_Tugas4_adi/handlers"
	"github.com/MIKTI_Tugas4_adi/middlewares"
	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	e := echo.New()

	e.POST("/login", handlers.Login(db))

	userGroup := e.Group("/users", middlewares.JWTMiddleware(), middlewares.IsAdmin)
	userGroup.POST("", handlers.CreateUser(db))

	todoGroup := e.Group("/todos", middlewares.JWTMiddleware(), middlewares.IsEditor)
	todoGroup.POST("", handlers.CreateTodoHandler(db))

	e.Logger.Fatal(e.Start(":8010"))
}
