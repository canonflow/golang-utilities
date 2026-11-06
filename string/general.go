package string

import (
	"encoding/json/v2"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToString(v any) string {
	if v == nil {
		return ""
	}

	switch c := v.(type) {
	case string:
		return c
	case int:
		return strconv.Itoa(c)
	case int64:
		return strconv.FormatInt(c, 10)
	case uint64:
		return strconv.FormatUint(c, 10)
	case float64:
		return strconv.FormatFloat(c, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(c)
	default:
		return fmt.Sprintf("%v", c)
	}
}

func ToTitle(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	s = re.ReplaceAllString(s, " ")
	s = strings.Join(strings.Fields(s), " ")
	return cases.Title(language.English).String(strings.TrimSpace(s))
}

func ParseJSONArray(s string) ([]string, error) {
	// s is stored as a JSON array string, e.g. `["alice","bob","delta.com"]`

	var err error
	s = strings.TrimSpace(s)
	if s == "" {
		return []string{}, nil
	}

	var result []string
	if err = json.Unmarshal([]byte(s), &result); err != nil {
		return nil, err
	}

	return result, nil
}
