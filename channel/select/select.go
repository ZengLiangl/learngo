package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)

	var vales []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeVal int
		if len(vales) > 0 {
			activeWorker = w
			activeVal = vales[0]
		}
		select {
		case n := <-c1:
			vales = append(vales, n)
		case n := <-c2:
			vales = append(vales, n)
		case activeWorker <- activeVal:
			vales = vales[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println(vales)
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
