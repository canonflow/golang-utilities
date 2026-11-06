package slice

func Unique[T comparable](haystack []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(haystack))

	for _, v := range haystack {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
