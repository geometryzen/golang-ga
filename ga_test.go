package ga

import (
	"testing"
)

func TestCanonicalReorderingSign(t *testing.T) {
	var a uint16 = 1 << 1
	var b uint16 = 1 << 2
	crs := canonicalReorderingSign(a, b)
	if crs != 1 {
		t.Errorf("canonicalReorderingSign(%08b, %08b) => %v", a, b, crs)
	}
	crs = canonicalReorderingSign(b, a)
	if crs != -1 {
		t.Errorf("canonicalReorderingSign(%08b, %08b) => %v", b, a, crs)
	}
}
