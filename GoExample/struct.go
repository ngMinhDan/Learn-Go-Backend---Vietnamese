package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) person {

	p := person{name: name, age: 42}
	fmt.Println(p)
	return p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Tom", 30})
	fmt.Println(person{"Ryan", 40})

	fmt.Println(person{name: "Ann", age: 40})

	jon1 := newPerson("Jon1")

	// Nếu jon2 thay đổi tuổi ở đây, chỉ có jon2 thay đổi, jon1 không thay đổi
	jon2 := jon1
	jon2.age = 100
	fmt.Println("Jon1 : ", jon1, " Jon2: ", jon2)

	// Nếu jon3 thay đổi tuổi ở đây, cả jon1, jon3 thay đổi
	jon3 := &jon1
	jon3.age = 100
	fmt.Println("Jon1 : ", jon1, " Jon3: ", jon3)

}
