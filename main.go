package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/wakatara/spotter/database"
	"github.com/wakatara/spotter/models"
	"github.com/wakatara/spotter/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	// app.Settings.Templates = pug.New("./views", ".pug")
	initDatabase()
	defer database.DBConn.Close()

	router.SetupRoutes(app)

	app.Listen(3000)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "spotter.db")
	if err != nil {
		panic("Failed to connect to DB. 😢")
	}
	fmt.Println("DB connect successful.")

	database.DBConn.AutoMigrate(&models.Sighting{})
	database.DBConn.AutoMigrate(&models.Venue{})
	fmt.Println("Database Spotter table migrated")
}
