package main

import (
	
	"fmt"

	"log"

	"net/http"

)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, Docker with Go!")

}

func main() {

	http.HandleFunc("/", handler)

	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
