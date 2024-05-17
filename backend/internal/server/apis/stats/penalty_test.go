package stats

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// https://dev.to/koddr/go-fiber-by-examples-testing-the-application-1ldf
func TestPenaltyTypeHandler(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/api/v2/penaltyTypes",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/api/v2/penaltyType",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	penaltyTypes := map[string][]string{"penaltyTypes": {"Boarding", "Charging", "Slashing"}}
	data, err := json.Marshal(penaltyTypes)
	if err != nil {
		assert.FailNow(t, "Failed to marshal data")
	}

	app.Get("/api/v2/penaltyTypes", func(c *fiber.Ctx) error {
		c.Type("json")
		return c.Send(data)
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
