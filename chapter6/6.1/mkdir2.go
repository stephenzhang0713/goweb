package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("dir1/dir2/dir3", 0777)
	if err != nil {
		fmt.Println(err)
	}
}
