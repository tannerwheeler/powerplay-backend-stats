package models

type Logo struct {
	DbModel
	Image []byte `gorm:"type:bytea"` // No JSON because this is going to be sent back as an actual image
}
