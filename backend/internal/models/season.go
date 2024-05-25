package models

import "time"

type Season struct {
	DbModel
	End   time.Time `json:"end"`
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
}
