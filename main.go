package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Simple go web server")

	fileServer:=http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server starting at server 8080")

	err:=http.ListenAndServe(":8080", nil)
	if (err!=nil) {
		log.Fatal(err)
	}
}