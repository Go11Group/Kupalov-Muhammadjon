package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	url := "http://example.com"

	data := Payload{Name: "John", Age: 30}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response: %s\n", body)
}
