package main 

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("File Path: ")
	var path string
	fmt.Scanln(&path)

	Data , error := ioutil.ReadFile(path)
	ErrorHandler(error)

	fmt.Println("Here's Your File Data. The Requested File" , path)
	fmt.Println(string(Data))
}

func ErrorHandler(error error) {
	if error != nil {
		panic(error) // Use Panic To Kill The Proccess. Println To Show Only The Error
	}
}