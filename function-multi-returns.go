package main

import "fmt"

func vals() (int, int) {
	return 10, 12
}

func main() {
	a, b := vals()
	fmt.Println("a: ", a, "b: ", b)

	_, y := vals()
	fmt.Println("b: ", y)
}
