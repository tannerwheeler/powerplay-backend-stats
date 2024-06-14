package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s *dbTestingSuite) TestCreateRoster() {

	roster := models.Roster{}

	r, err := s.session.CreateRoster(&roster)
	s.Nil(err)

	s.Equal(uint(1), r.ID)

	rosters, err := s.session.GetRosters()
	s.Nil(err)

	s.Len(rosters, 1)
}
