package main

import (
	// "fmt"
	"fmt"
	"log"

	// "module/generator"
	// "module/generator"
	"module/model"

	// "module/model"
	"module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// u := postgres.NewUserRepo(db)
	// randomUsers := generator.GenerateUsers(30)
	// fmt.Println(randomUsers)
	// for i := 0; i < 30; i++ {
	// 	err := u.CreateUser(randomUsers[i])
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// f := model.FilterUser{}
	// id := 1
	// f.Id = &id
	// res, err := u.GetUser(f)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(res)
	
	p := postgres.NewProductRepo(db)
	// products := generator.GenerateProducts(30)
	// for i := 0; i < 30; i++ {
	// 	err := p.CreateProduct(products[i])
	// 	if err != nil {
	// 		log.Println("Cannot add product")
	// 	}
	// }
	pId := 2
	fp := model.FilterProducts{Id: &pId}
	res, err := p.GetProduct(fp)
	if err != nil {
		log.Println("Can not get product(s)")
		log.Println(err)
	}
	fmt.Println(res)
}
