package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Foo struct{}

func (Foo) Speak() string {
	return "Hello, I am Foo"
}

type Bar struct{}

func (*Bar) Speak() string {
	return "Hello, I am Bar"
}

func main() {
	var someSpeaker Speaker = Foo{}
	fmt.Println(someSpeaker.Speak())

	someSpeaker = &Bar{} // must have : Pointer
	fmt.Println(someSpeaker.Speak())
}
