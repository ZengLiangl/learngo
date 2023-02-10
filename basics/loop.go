package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		printFileContents(file)
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("forever")
	}
}

func main() {
	fmt.Println(convertToBin(13))
	printfile("basics/test.txt")

	s := `adbab
			sdsd,s-s""。aasdas
			sdds
			sdaa`
	printFileContents(strings.NewReader(s))
}
