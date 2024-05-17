package stats

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
}

func getPenaltyTypes(c *fiber.Ctx) error {

	penaltyTypes := stubPenaltyTypes()
	jsonData, err := json.Marshal(penaltyTypes)
	if err != nil {
		return err
	}

	c.Type("json")

	// Send JSON response
	return c.Send(jsonData)
}

func stubPenaltyTypes() map[string][]string {
	return map[string][]string{"penaltyTypes": []string{"Boarding", "Charging", "Slashing"}}
}
