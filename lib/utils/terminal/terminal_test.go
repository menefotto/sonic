package terminal

import "testing"

func TestGetDimensions(t *testing.T) {
	h, w := GetDimensions()
	t.Logf("Dimensions are %d %d\n", h, w)
}
