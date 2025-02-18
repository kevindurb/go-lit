package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kevindurb/go-lit/internal/app"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	app := app.App{DB: db}
	app.Run()
}
