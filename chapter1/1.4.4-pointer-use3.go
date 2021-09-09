package main

import "fmt"

func exchange(c, d *int) {
	t := *c
	*c = *d
	*d = t
}

func exchange2(c, d *int) {
	d, c = c, d
}

func main() {
	a, b := 6, 8
	x, y := 6, 8
	exchange(&a, &b)
	exchange2(&x, &y)
	fmt.Println(a, b)
	fmt.Println(x, y)
}
