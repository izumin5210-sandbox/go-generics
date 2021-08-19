package main

import "fmt"

func main() {
	fmt.Println(
		mapSlice([]int{1, 2, 3}, func(v int) int {
			return v * 2
		}),
	)

	fmt.Println(
		mapSlice([]string{"foobar", "baz"}, func(v string) int {
			return len(v)
		}),
	)
}
