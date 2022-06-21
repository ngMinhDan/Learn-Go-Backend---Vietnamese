
package main 

import "fmt"

func main() {
// panic("There is a error here !")
 // ham panic se dung hoan toan chuong trinh
 defer func() {
 s := recover()
 fmt.Println(s)
 }()
 
 panic("hello panic")

}
