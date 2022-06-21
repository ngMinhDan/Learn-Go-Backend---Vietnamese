package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum : ", sum)

	for index, value := range nums {
		if index == 1 {
			fmt.Println(value)
		}
	}

	kvs := map[string]string{"A": "Apple", "B": "Ball"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, string(c))
	}
}
