package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Create("./test_rename.txt")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir("test_rename", 0777)
	err = os.Rename("./test_rename.txt", "./test_rename/test_rename_new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
