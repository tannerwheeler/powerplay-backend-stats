package db

import (
	"time"

	"github.com/jak103/powerplay/internal/models"
)

func (s *dbTestingSuite) TestCreateRoster() {
	layout := "2006-01-02 15:04:05"
	value := "2023-12-14 17:09:47"

	date, _ := time.Parse(layout, value)

	captain = models.User{
		FirstName:    "John",
		LastName:     "Smith",
		Email:        "test@email.com",
		Password:     "password",
		Phone:        "7",
		Role:         nil,
		SkillLevel:   3,
		CurrentTeams: nil,
		DateOfBirth:  date,
	}

	dbCapt, err := s.session.CreateUser(captain)
	s.Nil(err)
	s.NotNil(dbCapt)

	capt, err := s.session.GetUserbyID(dbCapt.ID)
	s.Nil(err)
	s.NotNil(capt)

	user1 = models.User{
		FirstName:    "Jack",
		LastName:     "Smith",
		Email:        "test1@email.com",
		Password:     "password",
		Phone:        "8",
		Role:         nil,
		SkillLevel:   3,
		CurrentTeams: nil,
		DateOfBirth:  date,
	}

	user2 = models.User{
		FirstName:    "Jill",
		LastName:     "Smith",
		Email:        "test2@email.com",
		Password:     "password",
		Phone:        "9",
		Role:         nil,
		SkillLevel:   3,
		CurrentTeams: nil,
		DateOfBirth:  date,
	}

	dbUser1, err := s.session.CreateUser(user1)
	s.Nil(err)
	s.NotNil(dbUser1)

	dbUser2, err := s.session.CreateUser(user2)
	s.Nil(err)
	s.NotNil(dbUser2)

	playerIDs := [2]uint{}

	players, err := s.session.GetUsersByIDs(playerIDs)
	s.Nil(err)
	s.NotNil(players)

	roster := models.Roster{
		CaptainID: capt.ID,
		Captain:   *capt,
		PlayerIDs: players,
	}

	r, err := s.session.CreateRoster(&roster)
	s.Nil(err)
	s.notNil(r)

	s.Equal(uint(1), r.ID)

	rosters, err := s.session.GetRosters()
	s.Nil(err)

	s.Len(rosters, 1)
}
