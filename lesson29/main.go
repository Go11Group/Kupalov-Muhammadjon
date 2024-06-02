package main

import (
	// "log"

	// "module/generator"

	// "module/model"
	"module/model"
	"module/printer"
	"module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	// ketma ket komentdan chiqalib ishlaturasiz bo'lmasa array bo'shab qolib error berishi mumkin
	u := postgres.NewUserRepo(db)
	// users := generator.GenerateRandomUser()
	// for i := 0; i < 30; i++ {
	// 	u.Create(users[i])
	// }

	// users, err := u.GetAllUsers()
	// if err != nil {
	// 	log.Println("Cannot get users")
	// }
	// printer.PrintUsers(users)
	filter := model.Filter{Age: 20}
	users, err := u.GetByFilter(filter)
	if err != nil {
		panic(err)
	}
	printer.PrintUsers(users)
	// user, err := u.GetById((*users)[0].ID)
	// if err != nil {
	// 	log.Println("Cannot get users by id")
	// }
	// printer.PrintUser(*user)

	// users, err = u.GetByFirstName("Alice")
	// if err != nil {
	// 	log.Println("Cannot get users by first_name")
	// }
	// printer.PrintUsers(users) = "Chello"
	// err = u.UpdateUser((*users)[0])
	// if err != nil {
	// 	log.Println("Cannot update user")
	// }
	// printer.PrintUser((*users)[0])
	// err = u.UpdateUser((*users)[0])
	// if err != nil {
	// 	log.Println("Cannot update user")
	// }
	// printer.PrintUser((*users)[0])

	// err = u.DeleteUser((*users)[0].ID)
	// if err != nil {
	// 	log.Println("Cannot delete user")
	// }
}
