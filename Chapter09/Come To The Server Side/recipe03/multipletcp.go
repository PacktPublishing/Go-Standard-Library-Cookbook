package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	ID := 0
	for {
		fmt.Println("Waiting for client...")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Client ID: %d connected.\n", ID)
		go func(c net.Conn, clientID int) {
			fmt.Fprintf(conn, "Welcome client ID: %d \n", clientID)
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					break
				}
				_, err = io.WriteString(conn, "Received: "+string(msg))
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			fmt.Println("Closing connection")
			conn.Close()
		}(conn, ID)
		ID++
	}

}
