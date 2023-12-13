package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Welcome to day 9")
	f := ParseFileToStringByNewLine("input.txt")

	first := FirstStar(f)
	fmt.Printf("First Star Result: %v%v", first, "\r\n")

	second := SecondStar(f)
	fmt.Printf("Second Star Result: %v%v", second, "\r\n")
}

func FirstStar(f []string) (res int) {
	parsedInput := [][]int{}

	for _, v := range f {
		splitStrs := regexp.MustCompile("\\s").Split(v, -1)
		splitNums := []int{}
		for _, w := range splitStrs {
			if num, err := strconv.Atoi(w); err == nil {
				splitNums = append(splitNums, num)
			}
		}
		parsedInput = append(parsedInput, splitNums)
	}
	for _, v := range parsedInput {
		res += BruteForceIt(v)
	}
	return
}

func SecondStar(f []string) (res int) {
	parsedInput := [][]int{}

	for _, v := range f {
		splitStrs := regexp.MustCompile("\\s").Split(v, -1)
		splitNums := []int{}
		for _, w := range splitStrs {
			if num, err := strconv.Atoi(w); err == nil {
				splitNums = append(splitNums, num)
			}
		}
		parsedInput = append(parsedInput, splitNums)
	}
	for _, v := range parsedInput {
		res += BruteForceItBackwards(v)
	}
	return
}

func BruteForceIt(input []int) (res int) {
	tree := [][]int{input}
	zeroCount := 0
	
	for i, j := 0,0; zeroCount < len(tree[j]); i++ {
		if len(tree)-1 < j+1 && len(tree[j]) > 1{
			tree = append(tree, []int{})
		}
		if i < len(tree[j])-1 {
			difference := tree[j][i+1] - tree[j][i]
			tree[j+1] = append(tree[j+1], difference)
			if difference == 0 {
				zeroCount++
				if zeroCount == len(tree[j])-2 {
					break
				}
			}
		}
		if i == len(tree[j])-1 {
			zeroCount = 0
			i = -1
			j++
		 }
	}
	
	for i := len(tree)-1; i >= 0; i--{
		res += tree[i][len(tree[i])-1]
	}
	return
}

func BruteForceItBackwards(input []int) (res int) {
	tree := [][]int{[]int{}}
	for i := len(input)-1; i >= 0; i-- {
		tree[0] = append(tree[0], input[i])
	}

	zeroCount := 0
	
	for i, j := 0,0; zeroCount < len(tree[j]); i++ {
		if len(tree)-1 < j+1 && len(tree[j]) > 1{
			tree = append(tree, []int{})
		}
		if i < len(tree[j])-1 {
			difference :=  tree[j][i] - tree[j][i+1]
			tree[j+1] = append(tree[j+1], difference)
			if difference == 0 {
				zeroCount++
				if zeroCount == len(tree[j])-2 {
					break
				}
			}
		}
		if i == len(tree[j])-1 {
			zeroCount = 0
			i = -1
			j++
		 }
	}
	
	for i := len(tree)-1; i >= 0; i--{
		res = tree[i][len(tree[i])-1] - res
	}
	return
}