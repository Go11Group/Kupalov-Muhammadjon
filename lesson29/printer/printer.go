package printer

import (
	"fmt"
	"module/model"
)

func PrintUser(user model.User){
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.Email)
	fmt.Println(user.Age)
	fmt.Println(user.Field)
	fmt.Println(user.Gender)
	fmt.Println(user.IsEmployee)
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println()
}
func PrintUsers(users *[]model.User){
	for _, user := range *users {
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println(user.FirstName)
		fmt.Println(user.LastName)
		fmt.Println(user.Email)
		fmt.Println(user.Age)
		fmt.Println(user.Field)
		fmt.Println(user.Gender)
		fmt.Println(user.IsEmployee)
		fmt.Println("----------------------------------------------------------------------------------")
	}
	fmt.Println()
}