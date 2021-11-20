package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "demo.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("打开文件出错:%v\n", err)
	}
	fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content))
}
