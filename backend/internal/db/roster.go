package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetRosters() ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)

	err := s.Preload("Players").Preload("Captain").Find(&rosters)

	return resultsOrError(rosters, err)
}

func (s session) CreateRoster(roster *models.Roster) (*models.Roster, error) {
	result := s.Create(roster)

	return resultOrError(roster, result)
}

func (s session) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}

	err := s.Find(&user, "ID = ?", id)

	return resultOrError(user, err)
}

func (s session) GetUsersByIDs(ids []uint) ([]*models.User, error) {
	users := make([]*models.User, 0)

	err := s.Where("ID IN (?)", ids).Find(&users)

	return resultsOrError(users, err)
}
