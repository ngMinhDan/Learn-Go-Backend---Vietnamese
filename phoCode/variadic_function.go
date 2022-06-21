// variadic function : args khong gioi han cac input vao cho ham
package main

import "fmt"

func add(args ...int) (int, int) {
 total := 0
 for _, v := range args {
  total += v
 }
 return len(args), total
}

func main(){
 fmt.Println(add(1,2,3,4,5))
}
