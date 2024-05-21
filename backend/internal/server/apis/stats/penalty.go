package stats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
}

func getPenaltyTypes(c *fiber.Ctx) error {

	log.Debug("Handling getting all penalty types")
	db := db.GetSession(c)
	penaltyTypes, err := db.GetPenaltyTypes()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all penalty types from the database")
		return err
	}

	return responder.OkWithData(c, penaltyTypes)
}
