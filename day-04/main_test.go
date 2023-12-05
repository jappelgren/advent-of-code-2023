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

func TestCleanInput(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := [][]string{
		{"41 48 83 86 17", "83 86 6 31 17 9 48 53"},
		{"13 32 20 16 61", "61 30 68 82 17 32 24 19"},
		{"1 21 53 59 44", "69 82 63 72 16 21 14 1"},
		{"41 92 73 84 69", "59 84 76 51 58 5 54 83"},
		{"87 83 26 28 32", "88 30 70 12 93 22 82 36"},
		{"31 18 13 56 72", "74 77 10 23 35 67 36 11"},
	}
	res := CleanInput(input)

	for i, v := range res {
		for j, w := range v {
			if w != want[i][j] {
				t.Fatalf(`Expected CleanInput(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
			}
		}
	}

}
