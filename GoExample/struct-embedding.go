package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprint("base with num = ", b.num)
}

type container struct {
	base // embeding in struct
	str  string
}

func main() {
	co := container{
		base: base{
			num: 100,
		},
		str: "Hello Embedding",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())
	fmt.Println("also describe:", co.describe())

	type describe interface {
		describe() string
	}

	var d = co
	fmt.Println("also 2 describe:", d.describe())

}
