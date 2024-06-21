package main

import "Go11Group/Kupalov-Muhammadjon/lesson42/AManWhoStandsInTheMiddle/api/handler"


func main(){
	h := handler.NewHandler()
	server := handler.CreateServer(h)
	server.ListenAndServe()
}