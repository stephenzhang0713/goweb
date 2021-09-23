package main

import (
	"fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("test_cookie")
	fmt.Printf("cookie:%#v, err:%v\n", c, err)

	cookie := &http.Cookie{
		Name:   "test_cookie",
		Value:  "krrsklHhefUUUFSSKLAkaLlJGGQEXZLJP",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 3600,
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", testHandle)
	http.ListenAndServe(":8085", nil)
}
