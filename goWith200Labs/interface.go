package main

import (
 "fmt"
)

type Speaker interface {
 Speak() string // ai co cai method nay thi deu duoc goi la Speaker
}

type Foo struct {}

func (Foo) Speak() string {
 return "Hello, I am Foo"
}

func main() {
 var someSpeaker Speaker = Foo{}
 fmt.Println(someSpeaker.Speak())
}
