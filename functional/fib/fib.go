package fib

import (
	"fmt"
	"io"
	"strings"
)

type IntGen func() int

// 1, 1, 2, 3, 5, 8, 13...
//
//	a, b
//	   a, b
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}
