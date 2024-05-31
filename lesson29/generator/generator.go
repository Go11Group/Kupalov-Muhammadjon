package generator

import (
	"fmt"
	"math/rand"
	"module/model"
)

func GenerateRandomUser() []model.User {
	firstNames := []string{"John", "Jane", "Alice", "Bob", "Charlie", "Daisy", "Eve", "Frank", "Grace", "Hank"}
	lastNames := []string{"Doe", "Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Martinez", "Lee", "Anderson"}
	fields := []string{"Engineering", "Marketing", "Sales", "Human Resources", "Finance", "IT", "Operations", "R&D", "Support", "Legal"}
	genders := []string{"Male", "Female"}

	users := []model.User{}
	for i := 0; i < 30; i++ {
		newUser := model.User{
			FirstName:  firstNames[rand.Intn(len(firstNames))],
			LastName:   lastNames[rand.Intn(len(lastNames))],
			Email:      fmt.Sprintf("user%d@example.com", i),
			Password:   "password",
			Age:        rand.Intn(50) + 20, // Random age between 20 and 70
			Field:      fields[rand.Intn(len(fields))],
			Gender:     genders[rand.Intn(len(genders))],
			IsEmployee: rand.Intn(2) == 1, // Randomly assign true or false
		}
		users = append(users, newUser)
	}

	return users
}
