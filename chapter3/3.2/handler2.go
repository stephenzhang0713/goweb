package main

import (
	"fmt"
	"log"
	"net/http"
)

type WelcomeHandler struct {
	Language string
}

func (h WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", h.Language)
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/cn", WelcomeHandler{Language: "欢迎一起学go web！"})
	mux.Handle("/en", WelcomeHandler{Language: "Welcome you,let's learn Go web!"})

	server := &http.Server{Handler: mux, Addr: ":8082"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
