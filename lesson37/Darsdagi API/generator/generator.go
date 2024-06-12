package generator

import (
	"log"
	"math/rand"
	"swagger/handler"
	"swagger/model"
	"time"

	"github.com/bxcodec/faker/v3"
)

func GenerateMockData(handler *handler.Handler) {
	rand.Seed(time.Now().UnixNano())
	genders := []string{"male", "female"}

	for i := 0; i < 50; i++ {
		user := model.User{
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			Age:        rand.Intn(100),                   
			Gender:     genders[rand.Intn(len(genders))], 
			Nation:     faker.DomainName(),
			Feild:      faker.AmountWithCurrency(),
			ParentName: faker.Name(),
			City:       faker.ChineseFirstName(),
		}

		err := handler.UserRepo.CreateUser(user)
		if err != nil {
			log.Printf("Error creating user: %v", err)
		}
	}
}
