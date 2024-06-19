package db

import (
	"fmt"
	"time"

	"github.com/jak103/powerplay/internal/models"
)

func (s *dbTestingSuite) CreateGameHelper() models.Game {
	// Create initial data
	season := models.Season{Name: "Season 1"}
	dbSeason, err := s.session.CreateSeason(&season)
	s.Nil(err)
	fmt.Println("Created Season ID:", dbSeason.ID)

	league := models.League{Name: "League 1", SeasonID: dbSeason.ID}
	dbLeague, err := s.session.CreateLeague(&league)
	s.Nil(err)
	fmt.Println("Created League ID:", dbLeague.ID)

	team1 := models.Team{Name: "Team 1", LeagueID: dbLeague.ID}
	team2 := models.Team{Name: "Team 2", LeagueID: dbLeague.ID}
	dbTeam1, err := s.session.CreateTeam(&team1)
	s.Nil(err)
	fmt.Println("Created Team 1 ID:", dbTeam1.ID)
	dbTeam2, err := s.session.CreateTeam(&team2)
	s.Nil(err)
	fmt.Println("Created Team 2 ID:", dbTeam2.ID)

	venue := models.Venue{Name: "Stadium 1"}
	dbVenue, err := s.session.SaveVenue(&venue)
	s.Nil(err)
	fmt.Println("Created Venue ID:", dbVenue.ID)

	// Create a game
	game := models.Game{
		SeasonID:           dbSeason.ID,
		VenueID:            dbVenue.ID,
		Status:             models.SCHEDULED,
		HomeTeamID:         dbTeam1.ID,
		AwayTeamID:         dbTeam2.ID,
		HomeTeamLockerRoom: "Locker Room 1",
		AwayTeamLockerRoom: "Locker Room 2",
		Start:              time.Now(),
	}
	err = s.session.CreateGame(&game)
	s.Nil(err)
	fmt.Println("Created Game ID:", game.ID)

	return game
}

func (s *dbTestingSuite) TestUpdateGameByID() {
	game := s.CreateGameHelper()

	updatedGame := models.Game{
		Status: models.IN_PROGRESS,
	}
	err := s.session.UpdateGame(fmt.Sprint(game.ID), updatedGame)
	s.Nil(err, "Expected no error updating game")

	dbGame, err := s.session.GetGameByID(fmt.Sprint(game.ID))
	s.Nil(err, "Expected no error getting game")
	s.Equal(models.IN_PROGRESS, dbGame.Status)
}

func (s *dbTestingSuite) TestUpdateInvalidID() {
	// Use the helper to create a game
	game := s.CreateGameHelper()

	// Correct data for updating a game
	updatedGame := models.Game{
		Status: models.IN_PROGRESS,
	}

	// Update with invalid ID
	invalidID := "9999"
	err := s.session.UpdateGame(invalidID, updatedGame)
	s.NotNil(err, "Expected an error when updating game with invalid ID")
	s.Contains(err.Error(), "record not found")

	// Check that the game with valid ID still exists and is unmodified
	dbGame, err := s.session.GetGameByID(fmt.Sprint(game.ID))
	s.Nil(err, "Expected no error getting game with valid ID")
	s.NotNil(dbGame)
	s.Equal(models.SCHEDULED, dbGame.Status, "Expected the game status to remain unchanged")
}

func (s *dbTestingSuite) TestGetGameByID() {
	// Use the helper to create a game
	game := s.CreateGameHelper()

	// Correct call for getting the game by ID
	dbGame, err := s.session.GetGameByID(fmt.Sprint(game.ID))
	s.Nil(err, "Expected no error getting game with valid ID")
	s.NotNil(dbGame)

	// Validate that properties in dbGame are what you expect them to be
	s.Equal(game.ID, dbGame.ID, "Expected game ID to match")
	s.Equal(game.SeasonID, dbGame.SeasonID, "Expected season ID to match")
	s.Equal(game.VenueID, dbGame.VenueID, "Expected venue ID to match")
	s.Equal(game.Status, dbGame.Status, "Expected status to match")
	s.Equal(game.HomeTeamID, dbGame.HomeTeamID, "Expected home team ID to match")
	s.Equal(game.AwayTeamID, dbGame.AwayTeamID, "Expected away team ID to match")
	s.Equal(game.HomeTeamLockerRoom, dbGame.HomeTeamLockerRoom, "Expected home team locker room to match")
	s.Equal(game.AwayTeamLockerRoom, dbGame.AwayTeamLockerRoom, "Expected away team locker room to match")
	s.WithinDuration(game.Start, dbGame.Start, time.Second, "Expected start time to be within one second")
}

func (s *dbTestingSuite) TestGetInvalidID() {
	invalidID := "9999"
	dbGame, err := s.session.GetGameByID(invalidID)
	s.NotNil(err)
	s.Nil(dbGame)
	s.Contains(err.Error(), "record not found")
}
