package main

import "os"

func main() {
	err := os.RemoveAll("./test_rename")
	if err != nil {
		return
	}
}
