package main

import "fmt"

func main() {
	s := make([]string, 10)
	fmt.Println("Empty: ", s)

	s[0] = "A"
	s[1] = "B"
	s[2] = "C"
	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[0])
	fmt.Println("Len: ", len(s))

	s = append(s, "D")
	s = append(s, "E", "F")

	fmt.Println("Add: ", s)

	s_copy := make([]string, len(s))
	copy(s_copy, s)
	fmt.Println("Copy: ", s_copy)

	l1 := s[2:5]
	l2 := s[:4]
	l3 := s[5:]
	fmt.Println("l1: ", l1)
	fmt.Println("l2: ", l2)
	fmt.Println("l3: ", l3)

	t := []string{"a", "b", "c"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}
