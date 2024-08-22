package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"sittellalab.com.au/internal/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	handlers.InitHandlers(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
