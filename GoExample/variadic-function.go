package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("Nums : ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3)
	sum(2, 4, 6)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum(nums...)

	nums2 := [5]int{3, 4, 6, 7, 8}
	sum(nums2[1:3]...)
}
