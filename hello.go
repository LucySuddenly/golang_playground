package main

import "fmt"

func main() {
	// a := []int{5, 4, 3, 2, 1}
	// a = append(a, 0)
	// object := make(map[string]int)
	// object["key1"] = 1
	// object["key2"] = 2
	// object["banana"] = 9999
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println(object)
	// result, err := sqrt(6)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	n := 10
	increment(&n)
	fmt.Println(n)
}

func increment(x *int) {
	*x++
}

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("No negative numbers")
// 	}
// 	return math.Sqrt(x), nil
// }
