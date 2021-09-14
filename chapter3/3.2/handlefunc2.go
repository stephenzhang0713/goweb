package main

import (
	"fmt"
	"log"
	http "net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi web")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)

	server := &http.Server{Addr: ":8081", Handler: mux}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
