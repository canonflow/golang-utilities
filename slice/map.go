package slice

func Map[T, R any](haystack []T, fn func(T) R) []R {
	r := make([]R, len(haystack))

	for i, v := range haystack {
		r[i] = fn(v)
	}

	return r
}

func Take[T any](items []T, count int) []T {
	if count <= 0 {
		return []T{}
	}

	if len(items) < count {
		return items
	}

	return items[:count]
}
