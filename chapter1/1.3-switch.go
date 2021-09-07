package main

import "fmt"

func main() {
	var a = "love"
	switch a {
	case "love":
		fmt.Println("love")
	case "programming":
		fmt.Println("programming")
	default:
		fmt.Println("none")
	}

	var language = "golang"
	switch language {
	case "golang", "java":
		fmt.Println("popular languages")
	}

	var r int = 6
	switch {
	case r > 1 && r < 10:
		fmt.Println(r)
	}
}
