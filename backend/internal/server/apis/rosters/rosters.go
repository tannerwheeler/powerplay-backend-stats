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
	apis.RegisterHandler(fiber.MethodPost, "/rosters", auth.Public, postRoster)
	apis.RegisterHandler(fiber.MethodGet, "/rosters", auth.Public, getRosters)
	// apis.RegisterHandler(fiber.MethodPost, "/rosters", auth.Public, postRosters)
}

func postRoster(c *fiber.Ctx) error {
	db := db.GetSession(c)

	captEmail := c.Query("capt_email")
	if captEmail == "" {
		return responder.BadRequest(c, "Empty Captain Email, Pass in valid value")
	}

	capt, err := db.GetUserByEmail(captEmail)
	if err != nil {
		return responder.InternalServerError(c, "Error getting captain in database")
	}

	if capt == nil {
		return responder.BadRequest(c, "No captain in database")
	}

	// rosterName := c.Query("roster_name")
	// if rosterName == "" {
	// 	return responder.BadRequest(c, "Empty Roster Name, Pass in valid value")
	// }

	// email := c.Query("email")
	// if email == "" {
	// 	return responder.BadRequest(c, "Empty Email, Pass in valid value")
	// }

	// user, err := db.GetUser(email)
	// if err != nil {
	// 	return responder.InternalServerError(c, "Unable to add user to roster")
	// }

	// err := db.PostUserToRoster(rosterName, user)
	// if err != nil {
	// 	log.WithErr(err).Alert("Failed to add user to roster")
	// 	return err
	// }

	roster := createRoster(capt)

	err = db.PostRoster(roster)
	if err != nil {
		return responder.InternalServerError(c, "Failed to create roster!")
	}

	return responder.Ok(c)
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

// func postRosters(c *fiber.Ctx) error {
// 	db := db.GetSession(c)
// 	err := db.PostRosters()
// 	if err != nil {
// 		log.WithErr(err).Alert("Failed to post roster to the database")
// 		return err
// 	}

// 	return responder.Ok(c)
// }

func createRoster(capt *models.User) *models.Roster {
	return &models.Roster{
		CaptainID: capt.ID,
		Captain:   *capt,
		Players:   nil,
	}
}
