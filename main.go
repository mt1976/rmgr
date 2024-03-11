package main

import (
	"fmt"
	"net"
)

func main() {

	port := 5382
	machine := "127.0.0.1"
	fmt.Println("Prototype Rate Manager Connector")
	fmt.Println("Machine  :", machine)
	fmt.Println("Port     :", port)
	fmt.Println("Method   :", "TCP")

	targetAddress := fmt.Sprintf("%s:%d", machine, port)

	listener, err := net.Listen("tcp", targetAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening for messages from", targetAddress)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s\n", buffer[:n])
	}
}
