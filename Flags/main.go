package main 

import (
	"fmt"
	"flag"
)

func main() {
	Name := flag.String("name" , "Mohammed" , "Your Name That Will Got Reflected On The CLI")
	Phone := flag.Int("phone" , 011 , "You Phone That Will Got Reflected On The CLI")

	flag.Parse()

	fmt.Println("Your name is:" , *Name)
	fmt.Println("Your phone is:" , *Phone)
}