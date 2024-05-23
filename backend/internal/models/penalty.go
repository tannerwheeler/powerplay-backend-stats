package models

type PenaltyType struct {
	DbModel
	Duration uint   `json:"duration"`
	Name     string `json:"name"`
	Severity string `json:"severity"`
}

type Penalty struct {
	DbModel
	CreatedBy     uint        `json:"created_by"`
	Duration      uint        `json:"duration"`
	GameID        uint        `json:"game_id"`
	PenaltyType   PenaltyType `json:"penalty_type"`
	PenaltyTypeID uint        `json:"penalty_type_id"`
	Period        uint        `json:"period"` // TODO: how to represent Shootouts?
	PlayerID      uint        `json:"player_id"`
	TeamID        uint        `json:"team_id"`
}
