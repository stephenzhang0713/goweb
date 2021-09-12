package main

import (
	"fmt"
	"goweb/chapter1/interface/oop1"
	"goweb/chapter1/interface/oop2"
)

type Num int

func (x Num) Equal(i int) bool {
	return int(x) == i
}

func (x Num) LessThan(i int) bool {
	return int(x) < i
}

func (x Num) BiggerThan(i int) bool {
	return int(x) > i
}

func main() {
	var f1 Num = 6
	var f2 oop1.NumInterface1 = f1
	var f3 oop2.NumInterface2 = f2
	fmt.Println(f2, f3)
}
