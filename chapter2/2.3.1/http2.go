package main

import (
	"log"
	"net/http"
)

func main() {
	//启动服务器
	srv := &http.Server{Addr: ":8088", Handler: http.HandlerFunc(handle)}
	log.Printf("Serving on https://0.0.0.0:8088")
	log.Fatal(srv.ListenAndServeTLS("chapter2/2.3.1/server.crt", "chapter2/2.3.1/server.key"))
}

//处理器函数
func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	w.Write([]byte("Hello this is a HTTP 2 message"))
}
