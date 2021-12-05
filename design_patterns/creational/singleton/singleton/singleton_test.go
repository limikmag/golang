package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instanceOne := GetInstance()

	if instanceOne == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	instanceTwo := GetInstance()
	if instanceOne != instanceTwo {
		t.Error("Expected same instance but it got a different instance")
	}

	counter := instanceTwo.AddOne()
	if counter != 1 {
		t.Error(fmt.Sprintf("Expected 1 but got %v", counter))
	}

	counter = instanceOne.AddOne()
	if counter != 2 {
		t.Error(fmt.Sprintf("Expected 2 but got %v", counter))
	}

}
