package main

import "os"

func main() {
	file, err := os.Create("WriteString.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	file.WriteString("Go web编程实战派从入门到精通")
}
