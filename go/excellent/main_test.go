package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		// This test checks if the function returns "even" for an even number HOGEHOGE
		t.Errorf("Expected 'even', got '%s'", result)
	}
}