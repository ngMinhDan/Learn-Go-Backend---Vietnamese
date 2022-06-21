package main 

import 
	(
		"fmt"
		"github.com/leekchan/accounting"
		"go-module/color"
		"go-module/client"
		"go-module/server"
	)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
    fmt.Println("Balance: ", ac.FormatMoney(123456789.213123))  // "$123,456,789.21"
    color.Println("Nguyen Minh Dan")


    fmt.Println("=============== CHAT SYSTEM ====================")
    for i := 0; i <100 ; i++ {
    	client.SendMessage("A", "Hello, How are you ?")
    	server.ReceiveMessage("B", "Oke, I am fine")
    	break
    }
}