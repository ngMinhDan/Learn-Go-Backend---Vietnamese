package main

import "fmt"

func main() {
	var a [5]int            // cách khai báo var + arrayName + [Độ dài]tên kiểu
	fmt.Println(a, "Empty") // Mặc định là 0

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("Get: ", a[4])

	fmt.Println("Len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	z := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = z
			z += 1
		}
	}
	fmt.Println(twoD)

}
