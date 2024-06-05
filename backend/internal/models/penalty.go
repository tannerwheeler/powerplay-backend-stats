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
	TeamID        *uint        `json:"team_id" gorm:"NOT NULL"`
	GameID        *uint        `json:"game_id" gorm:"NOT NULL"`
	Period        *uint        `json:"period" gorm:"NOT NULL"`
	Duration      *uint        `json:"duration" gorm:"NOT NULL"`
	CreatedBy     *uint        `json:"created_by" gorm:"NOT NULL"`
	PenaltyType   *PenaltyType `json:"penalty_type" gorm:"NOT NULL"`
	PenaltyTypeID *uint        `json:"penalty_type_id"`
}
