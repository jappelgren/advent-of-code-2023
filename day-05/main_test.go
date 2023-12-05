package main

import (
	"testing"
)

func TestFirstStar(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := 13
	res := FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}

func TestSecondStar(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := 30
	res := SecondStar(input)

	if res != want {
		t.Fatalf(`Expected SecondStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}