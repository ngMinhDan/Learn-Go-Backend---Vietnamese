package main

import "fmt"

func main() {
	product := make(map[string]interface{})
	product["name"] = "Iphone 13 Pro max"
	product["price"] = 200
	product["quantity"] = 100

	fmt.Println(product)

}
