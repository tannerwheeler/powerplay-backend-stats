package auth

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Role string

const (
	None        Role = "none"
	Player      Role = "player"
	Captain     Role = "captain"
	Referee     Role = "referee"
	ScoreKeeper Role = "scorekeeper"
	Manager     Role = "manager"
)

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		*r = ""
		return nil
	}
	val, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for Role")
	}
	*r = Role(val)
	return nil
}

func HasCorrectRole(usersRoles []Role, roles []Role) bool {
	for _, usersRole := range usersRoles {
		for _, neededRole := range roles {
			if neededRole == None || usersRole == neededRole {
				return true
			}
		}
	}

	return false
}

type Roles []Role

var (
	Public        Roles = []Role{None}
	Authenticated Roles = []Role{Manager, Referee, ScoreKeeper, Captain, Player}
	Staff         Roles = []Role{Manager, Referee, ScoreKeeper}
	ManagerOnly   Roles = []Role{Manager}
)

// Scan implements the Scanner interface for RoleSlice
func (rs *Roles) Scan(value interface{}) error {
	if value == nil {
		*rs = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for RoleSlice")
	}

	return json.Unmarshal(bytes, rs)
}

// Value implements the driver Valuer interface for RoleSlice
func (rs Roles) Value() (driver.Value, error) {
	return json.Marshal(rs)
}
