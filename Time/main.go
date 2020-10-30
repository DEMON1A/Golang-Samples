package main 

import (
	"fmt"
	"time"
)

func main() {
	Time := time.Now()
	fmt.Println("Current Time is:" , Time.String())
}