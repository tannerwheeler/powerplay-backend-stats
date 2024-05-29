// TODO: Create season post endpoint, as well as getters.

package season

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/season", auth.Public, getSeasonsHandler)
	// apis.RegisterHandler(fiber.MethodPost, "/season", auth.Public, getLeaguesHandler)

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


// func postSeasonsHandler(c *fiber.Ctx) error {
// 	retur nil
// }