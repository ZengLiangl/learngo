package main

import (
	"fmt"
)

func TryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("I don't what to do: %v", r))
		}
	}()
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic("1212")
	//panic(errors.New("this is a error"))
}

func main() {

	TryRecover()
}
