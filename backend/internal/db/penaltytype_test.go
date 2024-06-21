package db

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/formatters"
)

func (s *dbTestingSuite) TestGetPenaltyTypeById() {
	penaltyType := models.PenaltyType{
		Name:     "Slashing",
		Duration: 2,
		Severity: "Minor",
	}
	dbPenaltyType, err := s.session.CreatePenaltyType(&penaltyType)
	s.Nil(err)
	id := formatters.UintToString(dbPenaltyType.ID)

	result, err := s.session.GetPenaltyTypeByID(id)
	s.Nil(err)
	s.Equal(dbPenaltyType.ID, result.ID)
	s.Equal(penaltyType.Name, result.Name)
	s.Equal(penaltyType.Duration, result.Duration)
	s.Equal(penaltyType.Severity, result.Severity)
}

// CreatePenaltyType is already being tested in the other tests

func (s *dbTestingSuite) TestUpdatePenaltyType() {
	penaltyType := models.PenaltyType{
		Name:     "Slashing",
		Duration: 2,
		Severity: "Minor",
	}
	dbPenaltyType, err := s.session.CreatePenaltyType(&penaltyType)
	s.Nil(err)

	dbPenaltyType.Name = "Slashing2"
	dbPenaltyType.Duration = 4
	dbPenaltyType.Severity = "Major"

	result, err := s.session.UpdatePenaltyType(dbPenaltyType)
	s.Nil(err)
	s.Equal(dbPenaltyType.ID, result.ID)
	s.Equal("Slashing2", result.Name)
	s.Equal(uint(4), result.Duration)
	s.Equal("Major", result.Severity)
}

func (s *dbTestingSuite) TestDeletePenaltyType() {
	// Arrange
	penaltyType1 := models.PenaltyType{
		Name:     "Slashing",
		Duration: 2,
		Severity: "Minor",
	}
	dbPenaltyType1, err := s.session.CreatePenaltyType(&penaltyType1)
	s.Nil(err)

	penaltyType2 := models.PenaltyType{
		Name:     "Tripping",
		Duration: 3,
		Severity: "Minor",
	}
	dbPenaltyType2, err := s.session.CreatePenaltyType(&penaltyType2)
	s.Nil(err)

	getAllBefore, err := s.session.GetPenaltyTypes()
	s.Nil(err)
	s.Equal(2, len(getAllBefore))

	// Act
	err = s.session.DeletePenaltyType(dbPenaltyType1)
	s.Nil(err)

	// Assert
	getAllAfter, err := s.session.GetPenaltyTypes()
	s.Nil(err)
	s.Equal(1, len(getAllAfter))
	s.Equal(dbPenaltyType2.ID, getAllAfter[0].ID)
	s.Equal(penaltyType2.Name, getAllAfter[0].Name)
	s.Equal(penaltyType2.Duration, getAllAfter[0].Duration)
	s.Equal(penaltyType2.Severity, getAllAfter[0].Severity)
}
