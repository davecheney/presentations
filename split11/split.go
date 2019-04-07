package split

import "strings"

// tag::split[]
// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	if len(s) > 0 {
		result = append(result, s)
	}
	return result
}

// end::split[]
