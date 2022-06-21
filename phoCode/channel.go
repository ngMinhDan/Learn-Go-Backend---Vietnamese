package main

import (
 "fmt"
 "time"
)

func pinger(c chan string){
 for i := 0; ; i++ {
   c <- "ping pong" //cu the pingpong duoc truyen vao channel
 }
}

func printer(c chan string) {
 for {
  msg := <- c
  fmt.Println(msg) // cu in message tu pingpong ra thoi
  time.Sleep(time.Second *1)
 }	
}

func main() {
 var c chan string = make(chan string)
 go pinger(c)
 go printer(c)
 var input string
 fmt.Scanln(&input)
}
