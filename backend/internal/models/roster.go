package models

type Roster struct {
	DbModel
	Season   Season `json:"season"`
	SeasonID uint
	Team     Team `json:"team"`
	TeamID   uint
	Players  []*User `json:"players" gorm:"many2many:player_rosters"`
	Staff    []*User `json:"staff" gorm:"many2many:staff_rosters"`
}
