package main

import (
	"fmt"
	"goweb/chapter1/person"
)

func main() {
	s := new(person.Student)
	s.SetName("Shirdon")
	fmt.Println(s.GetName())
}
