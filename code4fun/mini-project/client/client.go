package client 

import 
	(
		"go-module/color"
	)

func SendMessage(sender string, message string) {
	color.Println(sender, message, 1)
}