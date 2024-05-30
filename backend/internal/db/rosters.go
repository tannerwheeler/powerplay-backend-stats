package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetRosters() ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)
	err := s.connection.Find(&rosters)
	return resultsOrError(rosters, err)
}
