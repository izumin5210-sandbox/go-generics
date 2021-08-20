package main

func testSliceUtils() {
	printSubSection("mapSlice", testMapSlice)
	printSubSection("filterSlice", testFilterSlice)
	printSubSection("filterMapSlice", testFilterMapSlice)
	printSubSection("findSlice", testFindSlice)
	printSubSection("containsSlice", testContainsSlice)
	printSubSection("compactSlice", testCompactSlice)
	printSubSection("uniqSlice", testUniqSlice)
	printSubSection("groupSliceBy", testGroupSliceBy)
	printSubSection("differenceSlice", testDifferenceSlice)
}

func testMapSlice() {
	snippet(
		func() interface{} {
			return mapSlice([]int{1, 2, 3, 4}, func(v int) int {
				return v * 2
			})
		},
	)
	snippet(
		func() interface{} {
			return mapSlice([]string{"foobar", "baz"}, func(v string) int {
				return len(v)
			})
		},
	)
}

func testFilterMapSlice() {
	snippet(
		func() interface{} {
			return filterMapSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int) (int, bool) {
				return v * 2, v%2 == 0
			})
		},
	)
}

func testFilterSlice() {
	snippet(
		func() interface{} {
			return filterSlice([]int{1, 2, 3, 4, 5, 6}, func(v int) bool {
				return v/2 == 0
			})
		},
	)

	snippet(
		func() interface{} {
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

			nonAdminUsers := filterNotSlice(
				users,
				func(user *User) bool { return user.Admin },
			)

			return mapSlice(
				nonAdminUsers,
				func(user *User) string { return user.Name },
			)
		},
	)
}

func testFindSlice() {
	snippet(
		func() interface{} {
			type User struct {
				Name string
				Age  int
			}
			teens, _ := findSlice([]*User{{Name: "Alice", Age: 18}, {Name: "Bob", Age: 27}, {Name: "Carol", Age: 24}}, func(u *User) bool {
				return u.Age >= 20
			})
			return teens
		},
	)
}

func testContainsSlice() {
	snippet(
		func() interface{} {
			return containsSlice([]int{1, 5, 6}, 4)
		},
	)
}

func testCompactSlice() {
	snippet(
		func() interface{} {
			return compactSlice([]string{"foo", "bar", "", "baz", ""})
		},
	)

	snippet(
		func() interface{} {
			type User struct {
				ID   int
				Name string
			}

			return mapSlice(
				compactSlice([]User{{ID: 1, Name: "Alice"}, {}, {ID: 2, Name: "Bob"}}),
				func(user User) string { return user.Name },
			)
		},
	)

	// NOTE: compiler panics
	// snippet(
	//   func() interface{} {
	//     type User struct {
	//       ID   int
	//       Name string
	//     }
	//     return mapSlice(
	//       compactSlice([]*User{{ID: 1, Name: "Alice"}, nil, {ID: 2, Name: "Bob"}}),
	//       func(user *User) string { return user.Name },
	//     )
	//   },
	// )
}

func testUniqSlice() {
	snippet(
		func() interface{} {
			return uniqSlice([]string{"a", "a", "b", "b", "c"})
		},
	)
}

func testGroupSliceBy() {
	snippet(
		func() interface{} {
			return groupSliceBy([]int{1, 2, 3, 4, 5, 6}, func(v int) int { return v % 3 })
		},
	)
}

func testDifferenceSlice() {
	snippet(
		func() interface{} {
			return differenceSlice([]int{1, 1, 2, 2, 3, 3, 4, 5}, []int{1, 2, 4})
		},
	)
}
