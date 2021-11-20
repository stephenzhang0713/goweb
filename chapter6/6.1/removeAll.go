package main

import (
	"log"
	"os"
)

func main() {
	os.MkdirAll("test1/test2/test3", 0777)
	err := os.RemoveAll("test1")
	if err != nil {
		log.Fatal(err)
	}
}
