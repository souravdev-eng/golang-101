package main

import "testing" // Go's built-in test tools; no extra dependency is needed.

// Test names start with Test and accept *testing.T.
func TestTotal(t *testing.T) {
	got := Total(3, 4) // Exercise the function as a caller would.
	want := 12         // Four items at 3 each should cost 12.
	if got != want {
		// Fatalf reports the mismatch and stops this test.
		t.Fatalf("Total(3, 4) = %d; want %d", got, want)
	}
}
