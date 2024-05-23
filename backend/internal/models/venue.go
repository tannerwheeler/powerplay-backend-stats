package models

import "github.com/lib/pq"

type Venue struct {
	DbModel
	Address     string         `json:"address"`
	LockerRooms pq.StringArray `json:"locker_rooms" gorm:"type:text[]"`
	Name        string         `json:"name"`
}
