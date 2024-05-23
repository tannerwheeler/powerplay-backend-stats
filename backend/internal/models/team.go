package models

type Team struct {
	DbModel
	Color    string    //`json:"color"`
	Leagues  []*League `json:"leagues" gorm:"many2many:team_leagues;"`
	LogoPath string    //`json:"logo_path"`
	Name     string    //`json:"name"`
}
