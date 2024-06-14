package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetPenaltyTypes() ([]models.PenaltyType, error) {
	penaltyTypes := make([]models.PenaltyType, 0)
	err := s.Find(&penaltyTypes)
	return resultsOrError(penaltyTypes, err)
}

func (s session) GetPenaltyTypeByID(id string) (*models.PenaltyType, error) {
	var penaltyType *models.PenaltyType
	err := s.First(&penaltyType, "id = ?", id)
	return resultOrError(penaltyType, err)
}

func (s session) CreatePenaltyType(request *models.PenaltyType) (*models.PenaltyType, error) {
	err := s.Create(request)
	return resultOrError(request, err)
}

func (s session) UpdatePenaltyType(request *models.PenaltyType) (*models.PenaltyType, error) {
	err := s.Save(request)
	return resultOrError(request, err)
}

func (s session) DeletePenaltyType(request *models.PenaltyType) error {
	return s.Delete(request).Error
}
