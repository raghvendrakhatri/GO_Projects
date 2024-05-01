package main

import (
	"fmt"
	"log"
	"net/http"

	httpHandler "basic_go_server/handlers"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", httpHandler.FormHandler)
	http.HandleFunc("/hello", httpHandler.HelloHandler)

	fmt.Println("starting server at port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
