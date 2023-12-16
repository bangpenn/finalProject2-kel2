package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/novita/finalproject2-revisi/cmd/app"
	"github.com/novita/finalproject2-revisi/infrastructure/database"
)

func init() {
	database.Init()
}

func main() {
	app.Start()
}
