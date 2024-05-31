package models

import "time"

type Season struct {
	DbModel
	Name          string         `json:"name"`
	Start         time.Time      `json:"start"`
	End           time.Time      `json:"end"`
	Registrations []Registration `json:"registrations"`
	Schedule      []Game         `json:"schedule"`
	Leagues       []League       `json:"leagues"`
}


// type Registration struct {
// 	DbModel
// 	SeasonID  uint `json:"season_id"`
// 	UserID    uint
// 	User      User       `json:"user"`
// 	Questions []Question `type:"questions"`
// }

// type Question struct {
// 	DbModel
// 	RegistrationID uint
// 	Text           string `json:"text"`
// 	Answer         string `json:"answer"`
// 	Render         string `json:"render"`
// }
