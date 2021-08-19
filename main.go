package main

import "fmt"

func main() {
	mapSliceTest()
	filterSliceTest()
	containsSliceTest()
	compactSliceTest()
	uniqSliceTest()
	differenceSliceTest()
}

func mapSliceTest() {
	fmt.Println("## mapSlice")
	fmt.Println(
		"double elements",
		mapSlice([]int{1, 2, 3}, func(v int) int {
			return v * 2
		}),
	)

	fmt.Println(
		"length of each strings",
		mapSlice([]string{"foobar", "baz"}, func(v string) int {
			return len(v)
		}),
	)

	fmt.Println()
}

func filterSliceTest() {
	fmt.Println("## fliterSlice")
	fmt.Println(
		"extract even numbers",
		filterSlice([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
			return v/2 == 0
		}),
	)

	type User struct {
		ID    int
		Name  string
		Admin bool
	}

	users := []*User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob", Admin: true},
		{ID: 3, Name: "Carol"},
	}
	adminUsers := filterSlice(
		users,
		func(user *User) bool { return user.Admin },
	)
	fmt.Println(
		"extract admin users ",
		mapSlice(
			adminUsers,
			func(user *User) string { return user.Name },
		),
	)

	nonAdminUsers := filterNotSlice(
		users,
		func(user *User) bool { return user.Admin },
	)
	fmt.Println(
		"extract non-admin users ",
		mapSlice(
			nonAdminUsers,
			func(user *User) string { return user.Name },
		),
	)

	fmt.Println()
}

func containsSliceTest() {
	fmt.Println("## containsSlice")
	fmt.Println(
		"contains 4",
		containsSlice([]int{1, 5, 6}, 4),
	)

	fmt.Println()
}

func compactSliceTest() {
	fmt.Println("## compactSlice")
	fmt.Println(
		"compact strings",
		compactSlice([]string{"foo", "bar", "", "baz", ""}),
	)

	type User struct {
		ID   int
		Name string
	}

	fmt.Println(
		"compact structs",
		mapSlice(
			compactSlice([]User{{ID: 1, Name: "Alice"}, {}, {ID: 2, Name: "Bob"}}),
			func(user User) string { return user.Name },
		),
	)

	// compiler panics
	// fmt.Println(
	//   "compact structs",
	//   mapSlice(
	//     compactSlice([]*User{{ID: 1, Name: "Alice"}, nil, {ID: 2, Name: "Bob"}}),
	//     func(user *User) string { return user.Name },
	//   ),
	// )

	fmt.Println()
}

func uniqSliceTest() {
	fmt.Println("## uniqSlice")
	fmt.Println(
		`uniq(["a", "a", "b", "b", "c"]) =`,
		uniqSlice([]string{"a", "a", "b", "b", "c"}),
	)

	fmt.Println()
}

func differenceSliceTest() {
	fmt.Println("## differenceSlice")
	fmt.Println(
		"[1, 1, 2, 2, 3, 3, 4, 5] - [1, 2, 4] =",
		differenceSlice([]int{1, 1, 2, 2, 3, 3, 4, 5}, []int{1, 2, 4}),
	)

	fmt.Println()
}
