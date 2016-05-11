package main

import (
	"fmt"
	"net"
	"net/textproto"
	"bufio"
	"os"
)

var (
	CONN_HOST string
	CONN_PORT string

	verbose bool = false
)

func listen(host, port string) {
	// Set vars
	CONN_HOST = host
	CONN_PORT = port

	// Listen for incoming connections.
	l, err := net.Listen("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Notify when connection is established
		fmt.Println("Incoming connection from", conn.RemoteAddr())

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Close the connection when you're done with it.
	defer conn.Close()

	// Make connection reader
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)


	
	// Read connection into data line by line
	data := ""
	for {
		line, err := tp.ReadLine()
		if err != nil {
			break
		}
		data += line + "\n"
	}

	// If verbose, print data
	if verbose {
		fmt.Println(data)
	}

	// Generate random folder name for paste
	tempDir := randString(8)
	// Write to file
	write(data, tempDir)

	// Send a response back to person contacting us.
	conn.Write([]byte("http://" + CONN_HOST + "/paste/" + tempDir + "\n"))	
}