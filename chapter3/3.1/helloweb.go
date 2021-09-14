package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go web")
}

func main() {
	http.HandleFunc("/hello", helloworld)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
