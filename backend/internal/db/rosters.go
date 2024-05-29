package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetRosters(query RosterQuery) ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)

	teamName := query.TeamName
	if teamName != nil {
		s.connection.Where("")
	}

	err := s.connection.Find(&rosters)
	return resultsOrError(rosters, err)
}
