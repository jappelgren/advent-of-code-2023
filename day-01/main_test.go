package main

import (
	"testing"
)

func TestFirstStar(t *testing.T) {
	input := "sample.txt"
	want := 142
	parsed := ParseFileToBytes(input)
	res, err := FirstStar(parsed)

	if res != want || err != nil {
		t.Fatalf(`FirstStar(%q) should return %#v, instead received %v`, input, want, res)
	}
}
