package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Hello World")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}

}
