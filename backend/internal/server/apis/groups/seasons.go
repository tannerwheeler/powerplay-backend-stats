
package season

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/jak103/powerplay/internal/models"


)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/seasons", auth.Public, getSeasonsHandler)
	apis.RegisterHandler(fiber.MethodPost, "/seasons", auth.Public, postSeasonsHandler)

}

func getSeasonsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	seasons, err := db.GetSeasons()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all seasons from the database")
		return err
	}

	return responder.OkWithData(c, seasons)
}





func postSeasonsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Debug("body: %q", c.Request().Body())
	seasonPostRequest := new(models.Season)
	err := c.BodyParser(seasonPostRequest)
	
	// If valid structure in post request, continue on
	if err != nil{
		log.WithErr(err).Error("Failed to parse Season POST request.")
		return err
	}

	// Connect to database and insert season
	db := db.GetSession(c)
	record, err := db.SaveSeason(seasonPostRequest)
	
	if err != nil{
		log.WithErr(err).Alert("Failed to parse season request payload")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c, "Could not post season into database")
	}
	
	return responder.Ok(c)
}