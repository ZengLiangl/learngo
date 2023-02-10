package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {

	fmt.Println("hello world")
	euler()
	triangle()
	constfun()
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	pow := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(pow)

}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(caleTriangle(a, b))
}

func caleTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func constfun() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}
