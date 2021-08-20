package main

func testMapUtils() {
	printSubSection("mapKeys / mapValues", testMapKeysValues)
	printSubSection("mapTransformKeys / mapTransformValues", testMapTransformKeysValues)
}

func testMapKeysValues() {
	snippet(
		func() interface{} {
			return mapKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3})
		},
	)
	snippet(
		func() interface{} {
			return mapValues(map[string]int{"foo": 1, "bar": 2, "baz": 3})
		},
	)
}

func testMapTransformKeysValues() {
	snippet(
		func() interface{} {
			return transformMapKeys(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string) string {
				return k + k
			})
		},
	)
	snippet(
		func() interface{} {
			return transformMapValues(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(v int) int {
				return v * 2
			})
		},
	)
}
