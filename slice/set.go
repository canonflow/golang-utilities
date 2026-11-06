package slice

// SymmetricDifference returns the items that appear in only one of a and b
// Items from a come first, them items from b, each in their original order.
// Duplicates are removed, e.g. ([1, 2, 3], [2, 3, 5]) -> [1, 5]
func SymmetricDifference[T comparable](a, b []T) []T {
	inA := make(map[T]struct{}, len(a))
	for _, v := range a {
		inA[v] = struct{}{}
	}

	inB := make(map[T]struct{}, len(b))
	for _, v := range b {
		inB[v] = struct{}{}
	}

	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, v := range a {
		if _, ok := inB[v]; ok {
			continue
		}

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	for _, v := range b {
		if _, ok := inA[v]; ok {
			continue
		}

		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// Diff returns the items in source that aren't in filter, in their original order.
// It only works in one direction. Diff([]int{1, 2, 3}, []int{2, 3, 5}) returns [1], and the 5 is ignored.
// If source contains the same value twice, Diff keeps both copies, unless the unique is set to true
func Diff[T comparable](source []T, filter []T, unique bool) []T {
	seen := make(map[T]struct{}, len(filter))
	result := make([]T, 0)

	for _, v := range filter {
		seen[v] = struct{}{}
	}

	for _, v := range source {
		if _, ok := seen[v]; !ok {
			result = append(result, v)
			if unique {
				seen[v] = struct{}{} // copies of v are now skipped
			}
		}
	}

	return result
}

// A-B
func LeftDiff[T comparable](left []T, right []T, unique bool) []T {
	return Diff(left, right, unique)
}

// B-A
func RightDiff[T comparable](left []T, right []T, unique bool) []T {
	return Diff(right, left, unique)
}
