// thuc te se co nhieu cong viec duoc thuc hien song song voi nhau ~ no goi la 
// concurrency : 
// DAU TIEN LA goRoutine
// goroutine la ham co the chay dong thoi voi cac ham khac, 
// de ham chay dong thoi thi ta them go truoc ham ~ 

package main
import "fmt"

func f(n int) {
 for i := 0; i < 10; i++ {
   fmt.Println(n, ":", i)	
 }
}

func main() {
 go f(99)
 var input string
 fmt.Scanln(&input)
}
