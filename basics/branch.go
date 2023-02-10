package main

import (
	"fmt"
	"os"
)

func main() {
	const filename = "test.txt"
	if contents, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
	fmt.Println(grade(100))
}

func grade(score int) string {

	p := ""
	switch {
	case score < 60:
		p = "F"
	case score < 70:
		p = "E"
	case score < 80:
		p = "D"
	case score < 85:
		p = "C"
	case score < 90:
		p = "B"
	case score <= 100:
		p = "A"
	default:
		panic("score is error")
	}
	return p
}
