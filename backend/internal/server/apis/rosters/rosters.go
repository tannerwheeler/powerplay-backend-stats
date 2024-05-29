package roster

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/rosters", auth.Public, getRosters)
}

func getRosters(c *fiber.Ctx) error {

	query := &models.RosterQuery{}

	c.QueryParser(query)

	db := db.GetSession(c)
	penaltyTypes, err := db.GetRosters(query)
	if err != nil {
		log.WithErr(err).Alert("Failed to get all penalty types from the database")
		return err
	}

	return responder.OkWithData(c, penaltyTypes)
}
