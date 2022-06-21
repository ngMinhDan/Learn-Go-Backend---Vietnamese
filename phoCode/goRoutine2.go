package main 

import "fmt"

func f(str string) {
 for i := 0; i < 10 ; i ++ {
  fmt.Println(str, ":", i)
 }
}

func main() {
 go f("hello goRoutine") 
 fmt.Println("End when main function end!")
}
