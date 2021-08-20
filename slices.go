package main

func mapSlice[In any, Out any](inputs []In, mapFunc func(input In) Out) []Out {
	outputs := make([]Out, len(inputs))
	for i, input := range inputs {
		outputs[i] = mapFunc(input)
	}
	return outputs
}

func filterSlice[In any](inputs []In, filterFunc func(input In) bool) []In {
	outputs := make([]In, 0, len(inputs))
	for _, input := range inputs {
		if filterFunc(input) {
			outputs = append(outputs, input)
		}
	}
	return outputs
}

func filterNotSlice[In any](inputs []In, filterFunc func(input In) bool) []In {
	return filterSlice(inputs, func(input In) bool { return !filterFunc(input) })
}

func findSlice[In any](inputs []In, findFunc func(input In) bool) (result In, ok bool) {
	for _, input := range inputs {
		if findFunc(input) {
			result = input
			return
		}
	}
	return
}

type comparableInterface[T any] interface {
	Compare(a, b T) string
}

func containsSlice[In comparable](inputs []In, want In) bool {
	for _, input := range inputs {
		if input == want {
			return true
		}
	}
	return false
}

func compactSlice[In comparable](inputs []In) []In {
	var zero In
	return filterNotSlice(inputs, func(input In) bool { return input == zero })
}

func uniqSlice[In comparable](inputs []In) []In {
	set := makeSet[In](len(inputs))
	return filterSlice(inputs, func(input In) bool {
		if set.Contains(input) {
			return false
		}
		set.Add(input)
		return true
	})
}

func groupSliceBy[In comparable, Key comparable](inputs []In, keyFunc func(input In) Key) map[Key][]In {
	result := make(map[Key][]In, len(inputs))
	for _, input := range inputs {
		key := keyFunc(input)
		result[key] = append(result[key], input)
	}
	return result
}

func differenceSlice[In comparable](inputs []In, values []In) []In {
	set := setFromSlice(values)
	return filterNotSlice(inputs, func(input In) bool {
		return set.Contains(input)
	})
}
