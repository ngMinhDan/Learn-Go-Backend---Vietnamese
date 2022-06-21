package client 

import 
	(
		"go-module/color"
	)

func ReceiveMessage(sender string, message string) {
	color.Println(sender, message, 0)
}