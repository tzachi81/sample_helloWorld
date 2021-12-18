package tests_test

import (
	"hello/greet"
	"testing"
)

func TestGreet(t *testing.T) {
	greetResult := greet.Greet()
	if greet.Greet() != "Hello user!" {
		t.Errorf("Greet() failed, got: %s, want: %s.", greetResult, "Hello user!")
	}
}
