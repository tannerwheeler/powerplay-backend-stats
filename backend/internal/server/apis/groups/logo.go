package groups

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/logo", auth.Public, handleLogoUpload)
	apis.RegisterHandler(fiber.MethodGet, "/logo/:id:int", auth.Public, handleGetLogoByID)
}

func handleLogoUpload(c *fiber.Ctx) error {
	log := locals.Logger(c)
	file, err := c.FormFile("image") // This is the key for the file in the form
	if err != nil {
		log.WithErr(err).Alert("Failed to upload logo")
		return responder.BadRequest(c, "Failed to upload logo")
	}

	fileData, err := file.Open()
	if err != nil {
		return responder.BadRequest(c, "Failed to open logo")
	}
	defer fileData.Close()

	imageBuffer := make([]byte, file.Size)
	if _, err = fileData.Read(imageBuffer); err != nil {
		return responder.BadRequest(c, "Failed to read logo")
	}

	db := db.GetSession(c)
	logo := models.Logo{
		Image: imageBuffer,
	}
	err = db.SaveLogo(&logo)
	if err != nil {
		log.WithErr(err).Alert("Failed to save logo")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}

func handleGetLogoByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.WithErr(err).Alert("Invalid ID")
		return responder.BadRequest(c, "Invalid ID")
	}

	log := locals.Logger(c)
	db := db.GetSession(c)
	logo, err := db.GetLogoByID(uint(id)) // Not sure why this required an additional cast
	if err != nil {
		log.WithErr(err).Alert("Failed to get the logo from the database")
		return responder.InternalServerError(c)
	}

	// Send JSON response
	return responder.OkWithData(c, logo)
}
