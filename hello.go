package main

import (
	"fmt"
)

func main() {
	a := []int{5, 4, 3, 2, 1}
	a = append(a, 0)
	fmt.Println(a)
}
