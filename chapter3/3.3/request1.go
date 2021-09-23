package main

import (
	"fmt"
	"net/url"
)

func main() {
	path := "http://localhost:8082/article?id=1"
	p, _ := url.Parse(path)
	fmt.Println(p.Host)
	fmt.Println(p.User)
	fmt.Println(p.RawQuery)
	fmt.Println(p.RequestURI())
}
