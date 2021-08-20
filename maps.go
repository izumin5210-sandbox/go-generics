package main

func mapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func mapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func transformMapKeys[K1, K2 comparable, V any](m map[K1]V, transform func(key K1) K2) map[K2]V {
	m2 := make(map[K2]V, len(m))
	for k, v := range m {
		m2[transform(k)] = v
	}
	return m2
}

func transformMapValues[K comparable, V1, V2 any](m map[K]V1, transform func(value V1) V2) map[K]V2 {
	m2 := make(map[K]V2, len(m))
	for k, v := range m {
		m2[k] = transform(v)
	}
	return m2
}
