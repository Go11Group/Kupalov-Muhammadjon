package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error in setting up listener", err)
		return
	}
	defer listener.Close()
	history := ""
	users := map[net.Conn]string{}
	fmt.Println("Server is listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error in accepting connection: ", err)
			return
		}
		go handleConnection(conn, &history, users)
	}

}

func handleConnection(conn net.Conn, history *string, users map[net.Conn]string) {
	defer conn.Close()
	defer func() {
		delete(users, conn)
	}()
	fmt.Println("Connected client Address: ", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error at reading massege: ", err)
			return
		}
		if len(message) == 1{
			continue
		}
		if _, ok := users[conn]; !ok {
			users[conn] = message[:len(message)-1]
			continue
		}
		if len(*history) > 0 {
			*history = (*history)[:len(*history)-16]
		}
		*history += users[conn] + ": " + message + "\nEnter message: " + string(byte(3))
		fmt.Println("Recived message: ", message)
		for con := range users {
			_, err = con.Write([]byte(*history))
			if err != nil {
				log.Println("Error at writing massege: ", err)
				return
			}
		}
	}
}
