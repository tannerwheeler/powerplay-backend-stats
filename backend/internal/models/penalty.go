package models

type PenaltyType struct {
	DbModel
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
	Severity string `json:"severity"`
}

// TODO used a pointer rather than uint, this allows the NOT NULL tag to behave as expected
// Need to see if this is okay with team. Another possible soln is using sql.nullInt or go.validate
type Penalty struct {
	DbModel
	PlayerID      uint        `json:"player_id" validate:"required"`
	TeamID        uint        `json:"team_id" validate:"required"`
	GameID        uint        `json:"game_id" validate:"required"`
	Period        uint        `json:"period" validate:"required"`
	Duration      uint        `json:"duration" validate:"required"`
	CreatedBy     uint        `json:"created_by" validate:"required"`
	PenaltyType   PenaltyType `json:"penalty_type" validate:"required"`
	PenaltyTypeID uint        `json:"penalty_type_id" validate:"required"`
}
