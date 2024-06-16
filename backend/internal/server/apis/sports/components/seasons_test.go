package components

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/stretchr/testify/assert"
)

func TestPostSeasonBadRequest(t *testing.T) {
	app := apis.CreateTestApp()

	req, _ := http.NewRequest(fiber.MethodPost, "/api/v1/seasons", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}
