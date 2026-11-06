package slice

import "strings"

func In[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func InString(haystack []string, needle string, strict bool) bool {
	fn := func(s string) string {
		return s
	}

	if !strict {
		fn = func(s string) string {
			return strings.ToLower(s)
		}
	}

	needle = fn(needle)

	for _, v := range haystack {
		if fn(v) == needle {
			return true
		}
	}

	return false
}
