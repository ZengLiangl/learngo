package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling funcation %s with agrs (%d,%d)\n", opName, a, b)
	return op(a, b)

}

func div(a, b int) (q, r int) {
	return a * b, a % b
}

func sum(numbers ...int) {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	fmt.Println(s)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}
func main() {
	fmt.Println(apply(pow, 1, 2))
	q, _ := div(3, 2)
	fmt.Println(q)
	sum(1, 2, 3, 4, 5, 5, 6, 6)
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(swap2(3, 5))
}
