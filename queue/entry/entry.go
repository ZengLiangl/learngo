package main

import (
	"fmt"
	"learngo/queue"
)

func main() {

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q)
	fmt.Println(q.Pop())
}
