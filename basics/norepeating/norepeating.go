package main

import (
	"fmt"
)

func main() {
	s := "我爱中国adsads我爱中国"
	len := lengthNonRepeatingSubStr(s)
	fmt.Println(len, string([]rune(s)[:len]))

}

func lengthNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		// 0-0+1 > 0
		// 1-0+1 > 1
		// 2-0+1 > 2
		// 3-1+1 > 3
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
