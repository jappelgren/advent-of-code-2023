package main

import "fmt"

func main() {
	f := ParseFileToBytes("input.txt")
	first, err := FirstStar(f)

	if err != nil {
		fmt.Printf(`An error occured... somehow: %v`, err)
		return
	}
	fmt.Println(first)

	second, err := SecondStar(f)
	if err != nil {
		fmt.Printf(`An error occured... somehow: %v`, err)
		return
	}
	fmt.Println(second)
}

func FirstStar(f []byte) (int, error) {
	res := 0
	numStore := []int{}

	for _, v := range f {
		if num := int(v - 48); num > -1 && num < 10 {
			numStore = append(numStore, num)
		}

		if v == 10 && len(numStore) > 0 {
			res += numStore[0]*10 + numStore[len(numStore)-1]
			numStore = []int{}
		}
	}

	return res, nil
}

func SecondStar(f []byte) (int, error) {
	res := 0
	numStore := []int{}

	for i := range f {
		var three string
		var four string
		var five string

		if remaining := len(f) - i; remaining >= 5 {
			five = string(f[i : i+5])
		}
		if remaining := len(f) - i; remaining >= 4 {
			four = string(f[i : i+4])
		}
		if remaining := len(f) - i; remaining >= 3 {
			three = string(f[i : i+3])
		}

		switch f[i] {
		case 'e':
			if five == "eight" {
				numStore = append(numStore, 8)
			}
		case 'f':
			if four == "four" {
				numStore = append(numStore, 4)
			}
			if four == "five" {
				numStore = append(numStore, 5)
			}
		case 'n':
			if four == "nine" {
				numStore = append(numStore, 9)
			}
		case 'o':
			if three == "one" {
				numStore = append(numStore, 1)
			}
		case 's':
			if three == "six" {
				numStore = append(numStore, 6)
			}
			if five == "seven" {
				numStore = append(numStore, 7)
			}
		case 't':
			if three == "two" {
				numStore = append(numStore, 2)
			}
			if five == "three" {
				numStore = append(numStore, 3)
			}
		case 'z':
			if four == "zero" {
				numStore = append(numStore, 0)
			}
		case '0':
			numStore = append(numStore, 0)
		case '1':
			numStore = append(numStore, 1)
		case '2':
			numStore = append(numStore, 2)
		case '3':
			numStore = append(numStore, 3)
		case '4':
			numStore = append(numStore, 4)
		case '5':
			numStore = append(numStore, 5)
		case '6':
			numStore = append(numStore, 6)
		case '7':
			numStore = append(numStore, 7)
		case '8':
			numStore = append(numStore, 8)
		case '9':
			numStore = append(numStore, 9)
		case 10:
			res += numStore[0]*10 + numStore[len(numStore)-1]
			numStore = []int{}
		}
	}
	fmt.Println(numStore)
	return res, nil
}
