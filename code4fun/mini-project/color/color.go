package color

import 
	(
	  "github.com/fatih/color"
	  "time"
	)

func Println(name string, message string, isSender int){
	if isSender == 1 { 
		c := color.New(color.FgCyan).Add(color.Underline)
		c.Printf(name, "\n message \t  ", time.Now())
	} else {
		c := color.New(color.FgRed).Add(color.Underline)
		c.Printf(name, "\n message \t ", time.Now())
	}
}