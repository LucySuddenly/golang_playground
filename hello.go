package main

import (
	"fmt"
)

func main() {
	// a := []int{5, 4, 3, 2, 1}
	// a = append(a, 0)
	object := make(map[string]int)
	object["key1"] = 1
	object["key2"] = 2
	object["banana"] = 9999

	fmt.Println(object)
}
