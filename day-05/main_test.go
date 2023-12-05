package main

import (
	"testing"
)

func TestFirstStar(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := uint64(35)
	res := FirstStar(input, InitSeedToSoil)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}

func TestSecondStar(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := uint64(0)
	res := SecondStar(input, InitSeedToSoilV2)

	if res != want {
		t.Fatalf(`Expected SecondStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}