package main

import (
	"fmt"
	"net"
)

func checkError(err *error) bool {
	if *err != nil {
		fmt.Println("Error:", (*err).Error())
		return true
	}
	return false
}

func handleClient(conn *net.Conn) {
	fmt.Println("Client connection established:",
		(*conn).RemoteAddr().String())

	buff := make([]byte, 1024)

	//received data from the client
	n, err := (*conn).Read(buff)

	var packet string

	//prepare some response
	if checkError(&err) {
		packet = "Failed to read client!"
	} else {
		packet = "Server received: " + string(buff[:n])
	}

	packet += "\n"

	buff = []byte(packet)

	//write response data to the client
	_, err = (*conn).Write(buff)

	checkError(&err)

	//shutdown the client connection
	err = (*conn).Close()

	checkError(&err)

	fmt.Println("Client connection terminated")
}

func main() {
	fmt.Println("Starting server...")

	protocol := "tcp"
	endPoint := "127.0.0.1:9050"

	//open up a specific port and start listening
	ln, err := net.Listen(protocol, endPoint)

	if checkError(&err) {
		return
	}

	fmt.Println("Server bound to", endPoint)

	fmt.Println("Awaiting client...")

	for {
		//accept new connections
		conn, err := ln.Accept()

		if checkError(&err) {
			return
		}

		//handle connections on a separate thread
		go handleClient(&conn)
	}
}
