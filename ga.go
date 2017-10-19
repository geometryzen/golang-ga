// Package ga provides a Geometric Algebra.
package ga

import "math/bits"

// canonicalReorderingSign determines the sign required to get a ^ b into canonical order.
func canonicalReorderingSign(a uint16, b uint16) float64 {
	var sum int
	a = a >> 1
	for a != 0 {
		sum = sum + bits.OnesCount16(a&b)
		a = a >> 1
	}
	if sum&1 == 0 {
		// even number of swaps
		return 1.0
	}
	// odd number of swaps
	return -1.0
}
