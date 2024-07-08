package math

import "testing"

func TestSum(t *testing.T) {
	sum := Sum([]int{13, 2, 2})
	if sum != 17 {
		t.Errorf("Expected 17, got %d", sum)
	}
}
