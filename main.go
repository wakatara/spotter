package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	sighting "github.com/wakatara/spotter/controllers"
	"github.com/wakatara/spotter/database"
	"github.com/wakatara/spotter/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	// app.Settings.Templates = pug.New("./views", ".pug")
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

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

func setupRoutes(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) { c.Render("index", fiber.Map{"Title": "Hello World"}) })
	app.Get("/api/v1/sighting", sighting.GetAllSightings)
	app.Get("/api/v1/sighting/:id", sighting.GetSighting)
	app.Post("/api/v1/sighting", sighting.NewSighting)
	// app.Put("/api/v1/sighting/:id", species.UpdateSighting)
	app.Delete("/api/v1/sighting/:id", sighting.DeleteSighting)
}
