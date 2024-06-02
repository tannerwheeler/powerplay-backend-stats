package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetRoster(rosterName string) (*models.Roster, error) {
	roster := &models.Roster{}

	result := s.connection.Where("roster_name = ?", rosterName).First(roster)

	return resultOrError(roster, result)
}

func (s session) GetRosters() ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)

	err := s.connection.Find(&rosters)

	return resultsOrError(rosters, err)
}

func (s session) PostRoster(roster *models.Roster) error {
	result := s.connection.Create(roster)

	return result.Error
}

// func (s session) PostUserToRoster(rosterName string, user models.User) error {
// 	roster, err := s.GetRoster(rosterName)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

func (s session) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	result := s.connection.Where("email = ?", email).First(user)

	return resultOrError(user, result)
}
