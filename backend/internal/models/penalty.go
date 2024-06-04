package models

import (
	"database/sql"
)


type PenaltyType struct {
	DbModel
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
	Severity string `json:"severity"`
}

// TODO used a pointer rather than uint, this allows the NOT NULL tag to behave as expected
// Need to see if this is okay with team. Another possible soln is using sql.nullInt
type Penalty struct {
	DbModel
	PlayerID      uint        `json:"player_id"`
	TeamID        uint        `json:"team_id"`
	GameID        uint        `json:"game_id"`
	Period        uint        `json:"period"`
	Duration      sql.NullInt64        `json:"duration" gorm:"NOT NULL"`
	CreatedBy     *uint        `json:"created_by" gorm:"NOT NULL"`
	PenaltyType   PenaltyType `json:"penalty_type"`
	PenaltyTypeID uint        `json:"penalty_type_id"`
}
