package stats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PenaltyType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
}

func getPenaltyTypes(c *fiber.Ctx) error {

	var penaltyType PenaltyType

	dsn := "host=localhost user=postgres password=password dbname=powerplay port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Cannot connect to db, %v", err)
	}

	result := db.First(&penaltyType)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Warn("Error occurred while querying database")
		return responder.InternalServerError(c)
	} else {
		if penaltyType.ID != 0 || penaltyType.Name != "" {
			return responder.OkWithData(c, penaltyType)
		} else {
			return responder.OkWithData(c, "No Data Found")
		}
	}

	// penaltyTypes := stubPenaltyTypes()
	// jsonData, err := json.Marshal(penaltyTypes)
	// if err != nil {
	// 	return err
	// }

	// c.Type("json")

	// // Send JSON response
	// return c.Send(jsonData)
}

// func stubPenaltyTypes() map[string][]string {
// 	return map[string][]string{"penaltyTypes": {"Boarding", "Charging", "Slashing"}}
// }
