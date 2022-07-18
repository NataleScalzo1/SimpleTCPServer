package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	dstream, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstream.Close()

	for {
		con, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return

		}
		go handleRequest(con)
	}
}

func handleRequest(conn net.Conn) {

	defer conn.Close()
	
	for {
		
		message, _ := bufio.NewReader(conn).ReadString('\n')
		
		fmt.Print("Message Received:", string(message))
		
		newmessage := strings.ToUpper(message)
		
		conn.Write([]byte(newmessage + "\n"))

	}

}
