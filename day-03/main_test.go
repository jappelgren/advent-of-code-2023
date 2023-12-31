package main

import (
	"testing"
)

func TestFirstStar(t *testing.T) {
	input := ParseFileToBytes("example.txt")
	want := 4361
	res := FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

	input = ParseFileToBytes("example2.txt")
	want = 413
	res = FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

	input = ParseFileToBytes("example3.txt")
	want = 925
	res = FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

	input = ParseFileToBytes("example4.txt")
	want = 40
	res = FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
	
	input = ParseFileToBytes("example5.txt")
	want = 62
	res = FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

	input = ParseFileToBytes("example6.txt")
	want = 156
	res = FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}

	input = ParseFileToBytes("example7.txt")
	want = 635
	res = FirstStar(input)

	if res != want {
		t.Fatalf(`Expected FirstStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}

func TestSecondStar(t *testing.T) {
	input := ParseFileToBytes("example.txt")
	want := 467835
	res := SecondStar(input)

	if res != want {
		t.Fatalf(`Expected SecondStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
	}
}

func TestFindAdjacentIndices(t *testing.T) {
	i := 14
	rowLen := 10
	arrLen := 100
	want := []int{4,5,15,25,24,23,13,3}
	res := FindAdjacentIndices(i, rowLen, arrLen)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 

	i = 4
	rowLen = 10
	want = []int{5,15,14,13,3}
	res = FindAdjacentIndices(i, rowLen, arrLen)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 

	i = 94
	rowLen = 10
	want = []int{84,85,95,93,83}
	res = FindAdjacentIndices(i, rowLen, arrLen)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 

	i = 23
	rowLen = 10
	want = []int{13,14,24,34,33,32,22,12}
	res = FindAdjacentIndices(i, rowLen, arrLen)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 

	i = 5
	rowLen = 10
	want = []int{6,4}
	res = FindAdjacentIndices(i, rowLen, 10)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 

	i = 10
	rowLen = 10
	want = []int{0,1,11,21,20}
	res = FindAdjacentIndices(i, rowLen, 100)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 

	i = 19
	rowLen = 10
	want = []int{9,29,28,18,8}
	res = FindAdjacentIndices(i, rowLen, 100)
	
	if len(res) != len(want) {
		t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
	}

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`Expected FindAdjacentIndicies(%v) to equal %v, instead got %v.%v`, i, want, res, "\r\n")
		}
	} 
}

// func TestSecondStar(t *testing.T) {
// 	input := ParseFileToStringByNewLine("example.txt")
// 	want := 2286
// 	res := SecondStar(input)

// 	if res != want {
// 		t.Fatalf(`Expected SecondStar(%v) to equal %v, instead got %v.%v`, input, want, res, "\r\n")
// 	}
// }