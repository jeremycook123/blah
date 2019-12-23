package util

import "testing"

func TestUpper(t *testing.T) {
	expected := "HELLO"
	if ret := Upper("hello"); ret != expected {
		t.Errorf("Upper() = %q, want %q", ret, expected)
	}
}
