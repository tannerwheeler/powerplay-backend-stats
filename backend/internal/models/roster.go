package models

type Roster struct {
	DbModel
	Players   []*User `json:"players" gorm:"many2many:player_rosters"`
	Captain   User    `json:"captain"`
	CaptainID uint    `json:"captain_id"`
}

type RosterQuery struct {
	TeamName string `query:"team_name"`
}
