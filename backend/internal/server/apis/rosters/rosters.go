package roster

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/rosters", auth.Public, getRosters)
}

func getRosters(c *fiber.Ctx) error {

	db := db.GetSession(c)
	rosters, err := db.GetRosters()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all rosters from the database")
		return err
	}

	return responder.OkWithData(c, rosters)
}
