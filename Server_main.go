package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chhandakguria/GoAPIs/Go_Server"
)

func main() {

	fileServer := http.FileServer(http.Dir("./Go_Server"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", Go_Server.Hellofunc)
	http.HandleFunc("/form", Go_Server.FormFunc)

	fmt.Println("Start the server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
