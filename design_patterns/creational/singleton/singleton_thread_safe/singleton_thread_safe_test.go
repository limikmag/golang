package singleton

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	s["this"] = "that"
	s2 := New()
	if s2["this"] != "that" {
		t.Error("Expected same instance but it got a different instance")
	}
}
