package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	fmt.Println("Server starts listening on port 12345")
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnnection(c)
	}
}

func handleServerConnnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Receive: ", msg)
	}
	c.Close()
}

func main() {
	server()
}
