package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			a[i]++
			runtime.Gosched()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
