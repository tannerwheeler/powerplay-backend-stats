package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveLogo(logo *models.Logo) (*models.Logo, error) {
	result := s.Create(logo)
	return resultOrError(logo, result)
}

func (s session) GetLogoByID(id string) (*models.Logo, error) {
	var logo models.Logo
	err := s.First(&logo, "id = ?", id)
	return resultOrError(&logo, err)
}
