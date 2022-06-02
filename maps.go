package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["k1"] = "hello"
	m["k2"] = "world"

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// key, value := m["k2"]
	fmt.Println(m["k2"])

	// fmt.Println("key: ", key)
	// fmt.Println("value: ", value)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
