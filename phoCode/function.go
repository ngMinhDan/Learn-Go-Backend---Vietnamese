package main

import "fmt"

func main(){
 xs := []float64{98, 93, 77, 82, 88}
 total := 0.0
 for _, v := range xs {
  total += v
 }
 fmt.Println(total)

}
