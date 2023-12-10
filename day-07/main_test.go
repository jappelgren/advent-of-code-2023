package main

import (
	"testing"
)

func TestFirstStar(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := 6440
	res := FirstStar(input, 1)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}

func TestSecondStar(t *testing.T) {
	input := ParseFileToStringByNewLine("example.txt")
	want := 5905
	res := SecondStar(input)

	if res != want {
		t.Fatalf(`Expected SecondStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	input = ParseFileToStringByNewLine("example2.txt")
	want = 112135
	res = SecondStar(input)

	if res != want {
		t.Fatalf(`Expected SecondStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}

func TestScoreHand(t *testing.T) {
	// Five of a kind
	input := [5]int{13,13,13,13,13}
	want := 6
	res := ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	// Four of a kind
	input = [5]int{13,13,13,13,2}
	want = 5
	res = ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	// Full House
	input = [5]int{13,13,13,2,2}
	want = 4
	res = ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	// Three of a kind
	input = [5]int{13,13,13,5,2}
	want = 3
	res = ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

	// Two pair
	input = [5]int{13,2,12,12,13}
	want = 2
	res = ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	// One pair
	input = [5]int{13,4,8,13,2}
	want = 1
	res = ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	// High card
	input = [5]int{13,10,8,3,2}
	want = 0
	res = ScoreHand(input)

	if res != want {
		t.Fatalf(`Expected Score(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

}

func TestWhichWins(t *testing.T) {
	inputOne := [5]int{13,12,13,12,2}
	inputTwo := [5]int{13,13,13,13,13}
	want := 1
	res := WhichWins(inputOne, inputTwo)

	if res != want {
		t.Fatalf(`Expected Score(%v, %v) to equal %v, instead got %v.%v`, inputOne, inputTwo, want, res, "\r\n")
	}
	
	inputOne = [5]int{3,2,4,6,5}
	inputTwo = [5]int{2,3,4,5,6}
	want = 0
	res = WhichWins(inputOne, inputTwo)

	if res != want {
		t.Fatalf(`Expected Score(%v, %v) to equal %v, instead got %v.%v`, inputOne, inputTwo, want, res, "\r\n")
	}
}