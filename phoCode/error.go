// trong cac ngon ngu khac, thi chung ta co Runtime Error, nghia la loi khi chay moi biet 
// Chi khi chay chuong trinh thi moi biet co loi do 
// Python, Java, PHP se try/catch de bat loi 
// Tuy nhien trong Go thi se gan vao mot bien xu ly loi ...

package main 

import (
	"errors"
	"fmt"
)
func divide(arg1, arg2 int) (int, error) {
	if arg2 == 0 {
		return -1, errors.New("Can't divide by 0")
	}
	return arg1 / arg2, nil
}

func main() {
	arg1 := 1
	arg2 := 0
	result, err := divide(arg1, arg2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}