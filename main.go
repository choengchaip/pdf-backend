package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pdf-backend/activitys"
	"pdf-backend/authentication"
	"pdf-backend/context"
	"pdf-backend/pdf"
	"pdf-backend/users"
)

func main() {
	ctx := context.NewContext()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	authentication.NewHttp(e, ctx)
	users.NewHttp(e, ctx)
	pdf.NewHttp(e, ctx)
	activitys.NewHttp(e, ctx)

	e.Logger.Fatal(e.Start(":8080"))
}
