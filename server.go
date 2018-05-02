// A simple concurrent ftp server

package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connectionType = "tcp"
	defaultPort    = ":8080"
)

func handleConnection(conn net.Conn) {
	// to do -- implement
}

func listenForConnections(port string) {
	ln, err := net.Listen(connectionType, port)
	if err != nil {
		fmt.Printf("Failed to listen for incoming connections")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Failed to accept incoming connections")
		}
		go handleConnection(conn)
	}
}

func main() {
	port := defaultPort
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	listenForConnections(port)
	//  to do -- implement
	// use "context" for cancellations
	// use file i/o and buffer for transmitting files over tcp
}
