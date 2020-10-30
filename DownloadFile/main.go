package main 

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	fmt.Println("URL: ")
	var URL string
	fmt.Scanln(&URL)

	fmt.Println("Filename: ")
	var path string
	fmt.Scanln(&path)

	error := DownloadFile(path, URL)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Done. The File Has Been Downloaded From THis URL:" , &URL)
}

func DownloadFile(path string , url string) error {
	Response , error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}

	defer Response.Body.Close()

	Output , error := os.Create(path)
	if error != nil {
		fmt.Println(error)
	}

	defer Output.Close()

	_ , err := io.Copy(Output , Response.Body)
	return err
}