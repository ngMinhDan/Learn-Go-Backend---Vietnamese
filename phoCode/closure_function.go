package main 

import "fmt"

func main() {
 add := func(args ...int) {
   fmt.Println(len(args))
 }
 add(1,2,3)
 
}
