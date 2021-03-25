package hw10

import "testing"

func TestCopy(t *testing.T) {
	err := Copy("from.txt", "to.txt", 0, 0)
	if err != nil {
		t.Fatalf("got an error: %v", err)
	}
}
