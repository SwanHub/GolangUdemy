package benchmarking

import "fmt"

// Greet takes in a name string and returns a salutation using that string
func Greet(s string) string {
	return fmt.Sprintf("welcome sir!, %v", s)
}
