package slice

import "strings"

func HasPrefixes(haystack []string, str string, strict bool) bool {
	fn := func(s string) string {
		return s
	}

	if !strict {
		fn = func(s string) string {
			return strings.ToLower(s)
		}
	}

	str = fn(str)

	for _, v := range haystack {
		if v == "" {
			continue
		}

		if strings.HasPrefix(str, fn(v)) {
			return true
		}
	}

	return false
}
