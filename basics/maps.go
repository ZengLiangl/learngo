package main

import "fmt"

func main() {

	m := map[string]string{
		"name":    "tom",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int
	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting val")
	course, ok := m["course"]
	fmt.Println(course, ok)
	if tom, ok := m["tom"]; ok {
		fmt.Println(tom)
	} else {
		fmt.Println("key does not exist")
	}
}
