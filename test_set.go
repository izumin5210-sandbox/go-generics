package main

func testSet() {
	printSubSection("setFromSlice", testSetFromSlice)
	printSubSection("makeSet", testMakeSet)
}

func testSetFromSlice() {
	snippet(
		func() interface{} {
			set := setFromSlice([]int{1, 2, 3})

			return []bool{
				set.Contains(2),
				set.Contains(4),
			}
		},
	)
}

func testMakeSet() {
	snippet(
		func() interface{} {
			set := makeSet[int](3) // capacity
			// set.Add(1, 2)
			set.Add(1)
			set.Add(2)
			set.Add(3)

			return []bool{
				set.Contains(2),
				set.Contains(4),
			}
		},
	)
}
