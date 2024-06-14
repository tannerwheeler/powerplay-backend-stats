package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s *dbTestingSuite) TestCreateRoster() {
	roster := models.Roster{}

	g, err := s.session.CreateRoster(&roster)
	s.Nil(err)

	s.Equal(uint(1), g.ID)

	goals, err := s.session.GetRosters()
	s.Nil(err)

	s.Len(goals, 1)
}
