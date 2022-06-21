package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println("Contains: ", s.Contains("Hello world", "world"))
	fmt.Println(1 == 1)
	fmt.Println('A' == 'a')
	fmt.Println("HasPrefix: ", s.HasPrefix("Hello world", "world"))
	fmt.Println("Count: ", s.Count("Hello world", "l"))
	fmt.Println("Index: ", s.Index("Hello world", "l"))
	fmt.Println(string(("Hello World")[0]))
}
