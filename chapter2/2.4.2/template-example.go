package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("chapter2/2.4.2/template_example.tmpl")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}
	name := "我爱go语言"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", HelloHandleFunc)
	http.ListenAndServe(":8086", nil)
}
