package main

import "fmt"

type Num int

func (x Num) Equal(i Num) bool {
	return x == i
}

func (x Num) LessThan(i Num) bool {
	return x < i
}

func (x Num) MoreThan(i Num) bool {
	return x > i
}

func (x *Num) Multiple(i Num) {
	*x = *x * i
}

func (x *Num) Divide(i Num) {
	*x = *x / i
}

type NumI interface {
	Equal(i Num) bool
	LessThan(i Num) bool
	MoreThan(i Num) bool
	Multiple(i Num)
	Divide(i Num)
}

func main() {
	var x Num = 8
	var y NumI = &x
	fmt.Println(y)

}
