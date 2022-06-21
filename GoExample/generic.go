package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k, _ := range m {
		fmt.Println("Key : ", k)
		r = append(r, k)
	}
	return r

}

type List(T any) struct {
	head, tail *element[T]
}

type 

func main() {
	var m = map[int]string{1: "minhdan", 2: "test", 3: "develop"}
	fmt.Println("Keys m: ", MapKeys(m))
}
