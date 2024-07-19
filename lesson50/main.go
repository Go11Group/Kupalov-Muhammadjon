package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

var username, password = "qwerty", "1234"

func main(){
	http.HandleFunc("/login", LoginMiddleware(user))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func LoginMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request)  {
		auth := r.Header.Get("Authorization")
		if auth == ""{
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || pair[0] != username || pair[1] != password{
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handlerFunc.ServeHTTP(w, r)
	}
}

func user(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("NIce"))
}