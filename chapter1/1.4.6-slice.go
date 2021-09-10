package main

import "fmt"

func main() {
	var sliceBuilder [20]int
	for i := 0; i < 20; i++ {
		sliceBuilder[i] = i + 1
	}
	fmt.Println(sliceBuilder[5:15])
	fmt.Println(sliceBuilder[15:])
	fmt.Println(sliceBuilder[:2])

	b := []int{6, 7, 8}
	fmt.Println(b[:])

	var sliceStr []string
	var sliceNum []int
	var emptySliceNum = []int{}
	fmt.Println(sliceStr, sliceNum, emptySliceNum)
	fmt.Println(len(sliceStr), len(sliceNum), len(emptySliceNum))
	fmt.Println(sliceStr == nil)
	fmt.Println(sliceNum == nil)
	fmt.Println(emptySliceNum == nil)

	slice1 := make([]int, 6)
	slice2 := make([]int, 6, 10)
	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), len(slice2))
}
