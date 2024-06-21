package db

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/formatters"
)

func (s *dbTestingSuite) TestGetPenaltyById() {
	// Arrange
	penaltyType1 := models.PenaltyType{
		Name:     "Slashing",
		Duration: 2,
		Severity: "Minor",
	}
	_, err := s.session.CreatePenaltyType(&penaltyType1)
	s.Nil(err)

	penalty := models.Penalty{
		PlayerID:      1,
		TeamID:        2,
		GameID:        3,
		Period:        4,
		Duration:      5,
		CreatedBy:     6,
		PenaltyTypeID: 1,
	}
	dbPenalty, err := s.session.CreatePenalty(&penalty)
	s.Nil(err)
	id := formatters.UintToString(dbPenalty.ID)

	// Act
	result, err := s.session.GetPenaltyByID(id)
	s.Nil(err)

	// Assert
	s.Equal(dbPenalty.ID, result.ID)
	s.Equal(penalty.PlayerID, result.PlayerID)
	s.Equal(penalty.TeamID, result.TeamID)
}

// CreatePenalty is already being tested in the other tests

func (s *dbTestingSuite) TestUpdatePenalty() {
	// Arrange
	penaltyType1 := models.PenaltyType{
		Name:     "Slashing",
		Duration: 2,
		Severity: "Minor",
	}
	_, err := s.session.CreatePenaltyType(&penaltyType1)
	s.Nil(err)

	penalty := models.Penalty{
		PlayerID:      1,
		TeamID:        2,
		GameID:        3,
		Period:        4,
		Duration:      5,
		CreatedBy:     6,
		PenaltyTypeID: 1,
	}
	dbPenalty, err := s.session.CreatePenalty(&penalty)
	s.Nil(err)

	// Act
	dbPenalty.PlayerID = 10
	dbPenalty.TeamID = 11
	dbPenalty.GameID = 12
	dbPenalty.Period = 13
	dbPenalty.Duration = 14
	dbPenalty.CreatedBy = 15
	dbPenalty.PenaltyTypeID = 1

	result, err := s.session.UpdatePenalty(dbPenalty)
	s.Nil(err)

	// Assert
	s.Equal(dbPenalty.ID, result.ID)
	s.Equal(uint(10), result.PlayerID)
	s.Equal(uint(11), result.TeamID)
	s.Equal(uint(12), result.GameID)
	s.Equal(uint(13), result.Period)
	s.Equal(uint(14), result.Duration)
	s.Equal(uint(15), result.CreatedBy)
	s.Equal(uint(1), result.PenaltyTypeID)
}

func (s *dbTestingSuite) TestDeletePenalty() {
	// Arrange
	penaltyType1 := models.PenaltyType{
		Name:     "Slashing",
		Duration: 2,
		Severity: "Minor",
	}
	_, err := s.session.CreatePenaltyType(&penaltyType1)
	s.Nil(err)

	penalty1 := models.Penalty{
		PlayerID:      1,
		TeamID:        2,
		GameID:        3,
		Period:        4,
		Duration:      5,
		CreatedBy:     6,
		PenaltyTypeID: 1,
	}
	dbPenalty1, err := s.session.CreatePenalty(&penalty1)
	s.Nil(err)

	penalty2 := models.Penalty{
		PlayerID:      10,
		TeamID:        11,
		GameID:        12,
		Period:        13,
		Duration:      14,
		CreatedBy:     15,
		PenaltyTypeID: 1,
	}
	dbPenalty2, err := s.session.CreatePenalty(&penalty2)
	s.Nil(err)

	getAllBefore, err := s.session.GetPenalties()
	s.Nil(err)
	s.Equal(2, len(getAllBefore))

	// Act
	err = s.session.DeletePenalty(dbPenalty1)
	s.Nil(err)

	// Assert
	getAllAfter, err := s.session.GetPenalties()
	s.Nil(err)
	s.Equal(1, len(getAllAfter))
	s.Equal(dbPenalty2.ID, getAllAfter[0].ID)
	s.Equal(penalty2.PlayerID, getAllAfter[0].PlayerID)
	s.Equal(penalty2.TeamID, getAllAfter[0].TeamID)
}
