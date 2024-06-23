package components

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/game", auth.Authenticated, handleCreateGame)
	apis.RegisterHandler(fiber.MethodPost, "/games", auth.Authenticated, handleCreateGames)
	apis.RegisterHandler(fiber.MethodGet, "/game/:id", auth.Authenticated, handleGetGameByID)
	apis.RegisterHandler(fiber.MethodGet, "/games", auth.Authenticated, handleGetGames)
	apis.RegisterHandler(fiber.MethodGet, "/games/:id", auth.Authenticated, handleGetGamesBySeason)
	apis.RegisterHandler(fiber.MethodPut, "/game/:id", auth.Authenticated, handleUpdateGameByID)
	apis.RegisterHandler(fiber.MethodPut, "/games", auth.Authenticated, handleUpdateGames)
	apis.RegisterHandler(fiber.MethodDelete, "/game/:id", auth.Authenticated, handleDeleteGameByID)
	apis.RegisterHandler(fiber.MethodDelete, "/games/:id", auth.Authenticated, handleDeleteGamesBySeason)
}

func handleCreateGame(c *fiber.Ctx) error {
	db := db.GetSession(c)
	var game models.Game
	if err := c.BodyParser(&game); err != nil {
		log.WithErr(err).Alert("Failed to parse game data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate request
	validate := validator.New()
	err := validate.Struct(game)
	if err != nil {
		return responder.BadRequest(c, "Failed to validate request")
	}

	if err := db.CreateGame(&game); err != nil {
		log.WithErr(err).Alert("Failed to create Game in the database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusCreated)
}

func handleCreateGames(c *fiber.Ctx) error {
	type Dto struct {
		Games []models.Game `json:"games"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	games := dto.Games
	session := db.GetSession(c)
	_, err = session.CreateGames(games)
	if err != nil {
		log.Error("Failed to save games to the database")
		return responder.InternalServerError(c, err)
	}
	return responder.CreatedWithData(c, games)
}

func handleGetGameByID(c *fiber.Ctx) error {
	type Dto struct {
		ID uint `json:"id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	id := dto.ID
	session := db.GetSession(c)
	game, err := session.GetGame(id)
	if err != nil {
		log.Error("Failed to get game from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.OkWithData(c, game)
}

func handleGetGames(c *fiber.Ctx) error {
	session := db.GetSession(c)
	games, err := session.GetGames()
	if err != nil {
		log.Error("Failed to get games from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.OkWithData(c, games)
}

func handleGetGamesBySeason(c *fiber.Ctx) error {
	type Dto struct {
		ID uint `json:"id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	id := dto.ID
	session := db.GetSession(c)
	games, err := session.GetGamesBySeason(id)
	if err != nil {
		log.Error("Failed to get games from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.OkWithData(c, games)
}

func handleUpdateGameByID(c *fiber.Ctx) error {
	type Dto struct {
		Id   uint        `json:"id"`
		Game models.Game `json:"game"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	game := dto.Game
	id := dto.Id
	session := db.GetSession(c)
	_, err = session.UpdateGame(id, game)
	if err != nil {
		log.Error("Failed to update game in the database")
		return responder.InternalServerError(c, err)
	}
	return responder.CreatedWithData(c, game)
}

func handleUpdateGames(c *fiber.Ctx) error {
	type Dto struct {
		Games []models.Game `json:"games"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	games := dto.Games
	session := db.GetSession(c)
	_, err = session.UpdateGames(games)
	if err != nil {
		log.Error("Failed to update games in the database")
		return responder.InternalServerError(c, err)
	}
	return responder.CreatedWithData(c, games)
}

func handleDeleteGameByID(c *fiber.Ctx) error {
	type Dto struct {
		ID uint `json:"id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	id := dto.ID
	session := db.GetSession(c)
	err = session.DeleteGame(id)
	if err != nil {
		log.Error("Failed to delete game from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c)
}

func handleDeleteGamesBySeason(c *fiber.Ctx) error {
	type Dto struct {
		SeasonID uint `json:"season_id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	seasonID := dto.SeasonID
	session := db.GetSession(c)
	err = session.DeleteGames(seasonID)
	if err != nil {
		log.Error("Failed to delete games from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c)
}
