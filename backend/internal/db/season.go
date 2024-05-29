package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetSeasons() ([]models.Season, error) {
	seasons := make([]models.Season, 0)
	err := s.connection.Find(&seasons)
	return resultsOrError(seasons, err)
}
