package main

import (
	"fmt"
	"net/http"
)

func getBody(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	http.HandleFunc("/getBody", getBody)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err)
	}
}
