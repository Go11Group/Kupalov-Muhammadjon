package mockdatagenerator

import (
	"internation/api/handler"
	"internation/model"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

func InsertUsers(handler *handler.Handler) {
	users := generateMockUsers(1000)
	for _, user := range users {
		handler.UserRepo.CreateUser(&user)
	}
}

func randomDate(start, end time.Time) time.Time {
	diff := end.Sub(start)
	sec := rand.Int63n(int64(diff.Seconds()))
	return start.Add(time.Duration(sec) * time.Second)
}

func generateMockUsers(n int) []model.User {
	var users []model.User

	start := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()

	for i := 0; i < n; i++ {
		user := model.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Birthday: randomDate(start, end),
			Password: faker.Password(),
		}
		users = append(users, user)
	}
	return users
}
