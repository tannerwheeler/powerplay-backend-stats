package models

type League struct {
	DbModel
	Name    string   //`json:"name"`
	Seasons []Season `json:"seasons" gorm:"many2many:league_season"`
}
