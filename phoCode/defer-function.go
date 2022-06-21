//defer co tac dung, chay mot lenh khac sau khi mot ham da ket thuc ~ callback
package main 

import "fmt"

func first(){
 fmt.Println("1st")
}

func second() {
 fmt.Println("2nd")
}

func main() {
 defer second() // chay sau khi ham main ket thuc 
 fmt.Println("Zero") // run first
 first() // 1
}

