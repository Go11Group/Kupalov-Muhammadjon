package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearTerminal()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error at connecting to server: ", err)
		return
	}
	defer conn.Close()
	go showResponse(conn)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ismingizni kiriting")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error at reading message: ", err)
			return
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Println("Error at sending message: ", err)
			return
		}
	}
}
func showResponse(conn net.Conn) {
	for {
		response, err := bufio.NewReader(conn).ReadString(byte(3))
		clearTerminal()
		if err != nil {
			log.Println("Error at reading response message: ", err)
			return
		}
		fmt.Println(response)
	}
}
