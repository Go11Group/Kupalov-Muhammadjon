package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	client := http.Client{}
	get(&client)
	post(&client)
	put(&client)
	delete(&client)

}


func get(client *http.Client) {
	req, err := http.NewRequest("GET", "http://localhost:8088/internation.uz/courses/all", nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bdy))
}
func post(client *http.Client) {
	body := `
	{
		"name": "Alice Johgfhnson",
		"email": "alice.johnhgfson@example.com",
		"birthday": "1990-01-01T00:00:00Z",
		"password": "passworbgfd123"
	}`
	preq, err := http.NewRequest("POST", "http://localhost:8088/internation.uz/users/create", bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}
	preq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(preq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	e, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(e))
}
func put(client *http.Client) {
	body := `
	{
		"name": "Alice Johgfhnson",
		"email": "alice.johnhgfson@example.com",
		"birthday": "1990-01-01T00:00:00Z",
		"password": "passworbgfd13"
  	}`
	ureq, err := http.NewRequest("PUT", "http://localhost:8088/internation.uz/users/89d8404a-86e6-4e6c-8f4c-ef489630ef84/update", bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}
	ureq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(ureq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	e, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(e))
}
func delete(client *http.Client) {
	dreq, err := http.NewRequest("DELETE", "http://localhost:8088/internation.uz/users/89d8404a-86e6-4e6c-8f4c-ef489630ef84/delete", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(dreq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	e, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(e))
}
