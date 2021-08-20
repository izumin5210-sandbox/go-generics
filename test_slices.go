package main

import "strings"

func testSliceUtils() {
	printSubSection("mapSlice", testMapSlice)
	printSubSection("flatMapSlice", testFlatMapSlice)
	printSubSection("filterSlice", testFilterSlice)
	printSubSection("filterMapSlice", testFilterMapSlice)
	printSubSection("findSlice", testFindSlice)
	printSubSection("indexSlice / lastIndexSlice", testIndexSlice)
	printSubSection("containsSlice", testContainsSlice)
	printSubSection("allSlice", testAllSlice)
	printSubSection("someSlice", testSomeSlice)
	printSubSection("maxSlice / minSlice", testMaxSliceMinSlice)
	printSubSection("compactSlice", testCompactSlice)
	printSubSection("uniqSlice", testUniqSlice)
	printSubSection("reduceSlice", testReduceSlice)
	printSubSection("groupSliceBy", testGroupSliceBy)
	printSubSection("reverseSlice", testReverseSlice)
	printSubSection("differenceSlice", testDifferenceSlice)
	printSubSection("intersectionSlice", testIntersectionSlice)
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

func testFlatMapSlice() {
	snippet(
		func() interface{} {
			return flatMapSlice([][]int{{1, 2}, {3, 4}}, func(v []int) []int {
				results := []int{}
				results = append(results, v...)
				results = append(results, 100)
				return results
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

func testIndexSlice() {
	snippet(
		func() interface{} {
			return indexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
				return strings.HasPrefix(v, "b")
			})
		},
	)

	snippet(
		func() interface{} {
			return indexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
				return strings.HasPrefix(v, "a")
			})
		},
	)

	snippet(
		func() interface{} {
			return lastIndexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
				return strings.HasPrefix(v, "b")
			})
		},
	)
	snippet(
		func() interface{} {
			return lastIndexSlice([]string{"foo", "bar", "baz"}, func(v string) bool {
				return strings.HasPrefix(v, "a")
			})
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

func testAllSlice() {
	snippet(
		func() interface{} {
			return allSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
				return len(v) >= 3
			})
		},
	)
	snippet(
		func() interface{} {
			return allSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
				return len(v) >= 4
			})
		},
	)
}

func testSomeSlice() {
	snippet(
		func() interface{} {
			return someSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
				return len(v) >= 4
			})
		},
	)
	snippet(
		func() interface{} {
			return someSlice([]string{"ant", "bear", "cat"}, func(v string) bool {
				return len(v) >= 5
			})
		},
	)
}

func testMaxSliceMinSlice() {
	snippet(
		func() interface{} {
			return maxSlice([]string{"albatross", "dog", "horse"}, func(v string) int {
				return len(v)
			})
		},
	)
	snippet(
		func() interface{} {
			return minSlice([]string{"albatross", "dog", "horse"}, func(v string) int {
				return len(v)
			})
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

func testReduceSlice() {
	snippet(
		func() interface{} {
			return reduceSlice([]int{1, 2, 3, 4, 5}, func(n int, acc int) int {
				return acc + n
			}, 0)
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

func testReverseSlice() {
	snippet(
		func() interface{} {
			return reverseSlice([]int{1, 2, 3, 4, 5})
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

func testIntersectionSlice() {
	snippet(
		func() interface{} {
			return intersectionSlice([]int{1, 2, 3}, []int{0, 1, 2})
		},
	)
}
