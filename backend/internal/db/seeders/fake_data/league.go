package fake_data

import (
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"gorm.io/gorm"
)

type LeagueSeeder struct{}

func (s LeagueSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	// Expecting the first argument to be the SeasonID
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	seasonID, ok := args[0].(uint)
	if !ok {
		return nil, errors.New("invalid type for SeasonID")
	}

	league := models.League{
		Name:          faker.Word(),
		CorrelationId: faker.UUIDHyphenated(),
		SeasonID:      seasonID,
		Teams:         []models.Team{},
	}
	if err := db.FirstOrCreate(&league, league).Error; err != nil {
		return nil, err
	}
	return league, nil
}
