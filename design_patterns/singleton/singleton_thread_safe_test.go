package singleton

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := singleton.New()
  s["this"] = "that"
  s2 := singleton.New()
  if s2["this"] != "that" {
    t.Error("Expected same instance but it got a different instance")
  }
}
