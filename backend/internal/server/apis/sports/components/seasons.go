package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/go-playground/validator/v10"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/seasons", auth.Public, getSeasonsHandler)
	apis.RegisterHandler(fiber.MethodPost, "/seasons", auth.Public, postSeasonsHandler)
}

func getSeasonsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := database.newSession(c)
	seasons, err := db.GetSeasons()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all seasons from the database")
		return err
	}

	return responder.OkWithData(c, seasons)
}

func postSeasonsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	seasonPostRequest := new(models.Season)
	err := c.BodyParser(seasonPostRequest)

	if err != nil {
		log.WithErr(err).Error("Failed to parse Season POST request.")
		return responder.BadRequest(c)
	}
	// Validate request
	validate := validator.New()
	err = validate.Struct(seasonPostRequest)
	if err != nil {
		return responder.BadRequest(c, "Failed to validate request")
	}

	db := database.newSession(c)
	record, err := db.CreateSeason(seasonPostRequest)

	if err != nil {
		log.WithErr(err).Alert("Failed to parse season request payload")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c, "Could not save season to database")
	}

	return responder.Ok(c)
}
