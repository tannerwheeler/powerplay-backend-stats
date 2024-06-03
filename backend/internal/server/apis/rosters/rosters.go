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
	// apis.RegisterHandler(fiber.MethodUpdate, "/rosters", auth.Public, updateRoster)
	// apis.RegisterHandler(fiber.MethodPost, "/rosters", auth.Public, postRosters)
}

// func updateRoster(c *fiber.Ctx) error {
// 	db := db.GetSession(c)

// 	captEmail := c.Query("capt_email")
// 	if captEmail == "" {
// 		return responder.BadRequest(c, "Empty Captain Email, Pass in valid value")
// 	}

// 	capt, err := db.GetUserByEmail(captEmail)
// 	if err != nil {
// 		return responder.InternalServerError(c, "Error getting captain in database")
// 	}

// 	if capt == nil {
// 		return responder.BadRequest(c, "No captain in database")
// 	}

// 	roster, err := db.GetRoster(capt.ID)
// 	if err != nil {
// 		return responder.BadRequest(c, "No roster in database")
// 	}

// 	userEmail := c.Query("user_email")
// 	if userEmail == "" {
// 		return responder.BadRequest(c, "Empty User Email, Pass in valid value")
// 	}

// 	user, err := db.GetUserByEmail(userEmail)
// 	if err != nil {
// 		return responder.InternalServerError(c, "Error getting user in database")
// 	}

// 	if user == nil {
// 		return responder.BadRequest(c, "No user in database")
// 	}

// 	players := roster.Players


// 	return responder.Ok(c)
// }

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

func createRoster(capt *models.User) *models.Roster {
	return &models.Roster{
		CaptainID: capt.ID,
		Captain:   *capt,
		Players:   nil,
	}
}