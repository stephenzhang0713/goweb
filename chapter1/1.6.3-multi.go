package main

import "fmt"

type Square struct {
	sideLen float32
}

type Triangle1 struct {
	Bottom float32
	Height float32
}

type Shape interface {
	Area1() float32
}

func (t *Triangle1) Area1() float32 {
	return (t.Bottom * t.Height) / 2
}

func (sq *Square) Area1() float32 {
	return sq.sideLen * sq.sideLen
}

func main() {
	t := &Triangle1{6, 8}
	s := &Square{8}
	shapes := []Shape{t, s}
	for n := range shapes {
		fmt.Println("图形数据是：", shapes[n])
		fmt.Println("它的面积是：", shapes[n].Area1())
	}
}
