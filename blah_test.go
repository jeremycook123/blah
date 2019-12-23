package blah

import "testing"

func TestBlah(t *testing.T) {
	expected := "Hello v1.0.0"
	if ret := Hello(); ret != expected {
		t.Errorf("Hello() = %q, want %q", ret, expected)
	}
}
