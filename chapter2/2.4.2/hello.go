package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("chapter2/2.4.2/hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := UserInfo{
		Name:   "李四",
		Gender: "男",
		Age:    28,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8080", nil)
}
