package slice

func Reduce[T, M any](haystack []T, fn func(M, T) M, initialValue M) M {
	acc := initialValue

	for _, v := range haystack {
		acc = fn(acc, v)
	}

	return acc
}
