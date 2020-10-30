package main 

import (
	"log"
	"io/ioutil"
	"net/http"
)

func main() {
	Response , error := http.Get("https://en6fu1edyzfl7.x.pipedream.net/")
	// fmt.Println(Response)
	if error != nil {
		log.Fatalln(error)
	}

	defer Response.Body.Close()
	Body , error := ioutil.ReadAll(Response.Body)
	if error != nil {
		log.Fatalln(error)
	}

	log.Println(string(Body))
	// log.Println(Response)
}