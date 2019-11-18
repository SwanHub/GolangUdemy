package bench2

import "strings"

// Cat concatenates a slice of strings
func Cat(xs []string) string {
	s := ""

	for _, v := range xs {
		s += v
		s += " "
	}
	return s
}

// Join function to bring a slice of strings together
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
