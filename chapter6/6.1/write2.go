package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("Go web编程实战派----从入门到精通")
	n, err := file.WriteAt([]byte("Go语言web"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
