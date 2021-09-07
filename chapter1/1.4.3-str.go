package main

import (
	"fmt"
)

func main() {
	str := "I love go web"
	for index, char := range str {
		fmt.Printf("%d %U %c \n", index, char, char)
	}
	chars := []rune(str)
	for _, char := range chars {
		fmt.Println(string(char))
	}

	str1 := "Hi 世界"
	by := []byte(str1)
	by[2] = ','
	fmt.Printf("%s\n", str1)
	fmt.Printf("%s\n", by)

	str2 := "Hi 世界"
	by1 := []rune(str2)
	by1[3] = '中'
	by1[4] = '国'
	fmt.Println(str2)
	fmt.Println(string(by1))

}
