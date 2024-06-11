package fake_data

import (
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s UserSeeder) Seed(db *gorm.DB, args ...interface{}) (interface{}, error) {
	roles := [][]auth.Role{auth.Authenticated, auth.Public, auth.Staff, auth.ManagerOnly}

	var createdUsers []models.User
	for i := 0; i < 4; i++ {
		randIndex := rand.Intn(len(roles))
		role := roles[randIndex]
		log.Debug("Role : %v, Type : %T", role, role)

		user := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			Password:     faker.Password(),
			Phone:        faker.Phonenumber(),
			Role:         role,
			SkillLevel:   rand.Intn(5),
			CurrentTeams: []models.Team{},
			DateOfBirth:  time.Time{},
		}

		log.Debug("User : %v, Type : %T", user, user)

		if err := db.FirstOrCreate(&user, models.User{}).Error; err != nil {
			return nil, err
		}

		createdUsers = append(createdUsers, user)
	}

	return createdUsers, nil
}
