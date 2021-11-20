package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.OpenFile("./demo.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件出错:%v\n", err)
		return
	}
	defer fp.Close()
}
