package main

import (
	"testing"
)

func TestParseFileToBytes (t *testing.T) {
	input := "txt-parser-test.txt"
	want := []byte{ 98, 105, 103, 32, 116, 105, 109, 101, 32, 116, 101, 115, 116, 32, 116, 101, 120, 116, 13, 10, 97, 32, 110, 101, 119, 32, 108, 
								  105,	110, 101, 32, 101, 118, 101, 110, 13, 10, 104, 111, 119, 32, 97, 98, 111, 117, 116, 32, 116, 104, 97, 116, 63, 13, 10 }
	res := ParseFileToBytes(input)

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`ParseFileToBytes(%v) should return %v, instead receieved %v`, input, want, res)
		}
	}
}

func TestParseFileToStringByNewLine(t *testing.T){
	input := "txt-parser-test.txt"
	want := []string{"big time test text", "a new line even", "how about that?"}
	res := ParseFileToStringByNewLine(input)

	for i, v := range res {
		if v != want[i] {
			t.Fatalf(`ParseFileToStringByNewLine(%v) should return %v, instead receieved %v`, input, want, res)
		}
	}
}