package router

import (
	"github.com/gofiber/fiber"
	sighting "github.com/wakatara/spotter/controllers"
)

func SetupRoutes(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) { c.Render("index", fiber.Map{"Title": "Hello World"}) })
	app.Get("/api/v1/sightings", sighting.GetAllSightings)
	app.Get("/api/v1/sightings/:id", sighting.GetSighting)
	app.Post("/api/v1/sightings", sighting.NewSighting)
	// app.Put("/api/v1/sighting/:id", sighting.UpdateSighting)
	app.Delete("/api/v1/sightings/:id", sighting.DeleteSighting)

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})
}
