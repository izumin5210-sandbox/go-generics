package main

func mapSlice[In any, Out any](inputs []In, mapFunc func(input In) Out) []Out {
	outputs := make([]Out, len(inputs))
	for i, input := range inputs {
		outputs[i] = mapFunc(input)
	}
	return outputs
}
