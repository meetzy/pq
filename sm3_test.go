package pq

import (
	"testing"
)

func TestSm3(t *testing.T) {
	if Sm3ToString("abc") != "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0" {
		t.Fatal("Sm3 error")
	}
}
