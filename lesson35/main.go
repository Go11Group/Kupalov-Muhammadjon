package main

import (
	// "fmt"
	// "leetcode/model"
	"leetcode/model"
	"leetcode/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	u := postgres.NewProblemRepo(db)
	newProblem := model.Problem{QuestionNumber: 1, Title: "Unique Path", DifficultyLevel: "Medium",
	Description: "Find pathes", Examples: []string{"like this"}, Hints: []string{"be calm"}}
	err = u.CreateProblem(newProblem)
	if err != nil {
		panic(err)
	}
	// f := model.UserFilter{}
	// users, err := u.GetProblems(f)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(users)
	// newUser := model.User{Id: "e50fe1db-8ca4-4918-8ad6-3ff665d5de3d", FullName: "Ko'palov Muhammadjon", Username: "kupalovmuhammadjon", Bio: "Go dev"}
	// err = u.UpdateProblem(newUser)
	// if err != nil {
	// 	panic(err)
	// }
	// err = u.DeleteProblem("e50fe1db-8ca4-4918-8ad6-3ff665d5de3d")
	// if err != nil {
	// 	panic(err)
	// }


}
