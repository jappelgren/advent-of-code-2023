package main

import (
	"fmt"
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
	fmt.Printf(`First Star test wanted %v, got %v. Test passed%v`, want, res, "\r\n")
}

func TestSecondStar(t *testing.T) {
	input := "sample2.txt"
	want := 281
	parsed := ParseFileToBytes(input)
	res, err := SecondStar(parsed)

	if res != want || err != nil {
		t.Fatalf(`SecondStar(%q) should return %#v, instead received %v`, input, want, res)
	}
	fmt.Printf(`Second Star test wanted %v, got %v. Test passed%v`, want, res, "\r\n")
}