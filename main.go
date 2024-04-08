package main

import (
	"goshop/database"
	"goshop/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	database.Connect()
	routes.Routes(e)
	e.Logger.Fatal(e.Start(":" + port))
}
