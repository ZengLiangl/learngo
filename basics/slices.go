package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100

}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr", arr)
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1", s1)
	s2 := arr[:]
	fmt.Println("s2", s2)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(s2)

	test()
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d \n", s, len(s), cap(s))
}

func test() {
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 3, 4, 6}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)

	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copy slice  ")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Printf("front:%v,tail%v \n ", front, tail)

	printSlice(s2)
}
