package sighting

import (
	"github.com/gofiber/fiber"
	"github.com/wakatara/spotter/database"
	"github.com/wakatara/spotter/models"
)

func GetAllSightings(c *fiber.Ctx) {
	db := database.DBConn
	var sightings []models.Sighting
	db.Find(&sightings)
	c.JSON(sightings)
}

func GetSighting(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var sighting models.Sighting
	db.Find(&sighting, id)
	c.JSON(sighting)
}

func NewSighting(c *fiber.Ctx) {
	db := database.DBConn

	sighting := new(models.Sighting)
	if err := c.BodyParser(sighting); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&sighting)
	c.JSON(sighting)
}

func DeleteSighting(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var sighting models.Sighting

	db.First(&sighting, id)
	if sighting.Registration == "" {
		c.Status(500).Send("No such sighting seems to exist with given ID.")
		return
	}
	db.Delete(&sighting)
	c.Send("Sighting successfully deleted.")
}
