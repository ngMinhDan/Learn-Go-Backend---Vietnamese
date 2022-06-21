package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msg := "Hello Server1"
		fmt.Println("Sending : ", msg)
		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 5)
	}
	// c.Close()

}

func main() {
	client()
}
