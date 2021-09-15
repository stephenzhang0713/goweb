package main

import (
	"fmt"
	"log"
	"net/http"
)

func hi1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi web")
}

func main() {
	http.HandleFunc("/", hi1)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
