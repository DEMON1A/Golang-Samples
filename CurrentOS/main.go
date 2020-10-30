package main 

import (
	"fmt"
	"runtime"
)

func main() {
	OS := runtime.GOOS
	fmt.Println("You Current OS is:" , OS)
}