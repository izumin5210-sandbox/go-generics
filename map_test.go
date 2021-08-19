package main

import (
	"fmt"
)

func Example_mapSlice() {
	outputs := mapSlice([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})

	// Output:
	// [2 4 6]
	fmt.Println(outputs)
}
