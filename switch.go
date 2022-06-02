package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Friday:
		fmt.Println("It is Weekend")
	default:
		fmt.Println("It is Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's is before noon")
	default:
		fmt.Println("It's afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's Boolean")
		case int:
			fmt.Println("It's Int")
		case string:
			fmt.Println("It's String")
		default:
			fmt.Println("Not define", t)
		}
	}

	whatAmI(true)
	whatAmI(100)
	whatAmI("hey yo")

}
