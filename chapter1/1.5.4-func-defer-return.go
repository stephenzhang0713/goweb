package main

import "fmt"

var name string = "go"

func myFunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc()函数里面的name: %s\n", name)
	return name
}

func main() {
	myname := myFunc()
	fmt.Printf("main()函数里面的name: %s\n", name)
	fmt.Println("main()函数里面的myname: ", myname)
}
