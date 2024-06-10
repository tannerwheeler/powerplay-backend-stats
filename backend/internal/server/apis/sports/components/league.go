package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/jak103/powerplay/internal/utils/validators"
	"strconv"
	"strings"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/leagues", auth.Public, getLeaguesHandler)
	apis.RegisterHandler(fiber.MethodPost, "/leagues", auth.Public, postLeagueHandler)
}

func getLeaguesHandler(c *fiber.Ctx) error {
	offsetParam := c.Query("offset", "0")
	limitParam := c.Query("limit", "10")
	fetchAll := c.Query("fetch_all", "false") // Default to false

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return responder.BadRequest(c, "Invalid offset parameter")
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return responder.BadRequest(c, "Invalid limit parameter")
	}

	fetchAllBool, err := strconv.ParseBool(fetchAll)
	if err != nil {
		return responder.BadRequest(c, "Invalid fetch_all parameter")
	}

	sortField := c.Query("sort_field", "ID")
	sortOrder := strings.ToUpper(c.Query("sort_order", "ASC"))
	if sortOrder != "ASC" && sortOrder != "DESC" {
		return responder.BadRequest(c, "Invalid sort_order parameter")
	}

	if !validators.IsValidSortField(sortField, models.League{}) {
		return responder.BadRequest(c, "Invalid sort_field parameter")
	}

	log := locals.Logger(c)
	db := db.GetSession(c)

	if fetchAllBool {
		leagues, err := db.GetLeagues(sortField, sortOrder)
		if err != nil {
			// todo: create ticket to standardize this error message and pass in model name
			log.WithErr(err).Alert("Failed to get all leagues from the database")
			return err
		}
		return responder.OkWithData(c, leagues)
	} else {
		leagues, err := db.GetLeaguesPaginated(offset, limit, sortField, sortOrder)
		if err != nil {
			log.WithErr(err).Alert("Failed to get leagues from the database")
			return responder.InternalServerError(c)
		}
		return responder.OkWithData(c, leagues)
	}

}

func postLeagueHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)

	leagueRequest := &models.League{}
	err := c.BodyParser(leagueRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse leagues request payload")
		return responder.BadRequest(c, "Failed to parse leagues request payload")
	}

	db := db.GetSession(c)
	err = db.CreateLeague(leagueRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to save leagues request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
