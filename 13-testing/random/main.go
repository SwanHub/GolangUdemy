// Package random is so random, it sums your integer values neatly
package random

// Sum2 sums all integers received as arguments
func Sum2(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
