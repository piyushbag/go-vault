package main

import (
	"fmt"

	"github.com/piyushbag/twg/math" // Add import statement
)

func main() {
	sum := math.Sum([]int{11, 4, 15})
	if sum != 30 {
		msg := fmt.Sprintf("FAIL: Expected 30, got %d", sum)
		panic(msg)
	}
	add := math.Add(1, 2)
	if add != 3 {
		msg := fmt.Sprintf("FAIL: Expected 3, got %d", add)
		panic(msg)
	}
	fmt.Println("PASS")
}
