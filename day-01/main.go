package main

import "fmt"

func main() {
	f := ParseFileToBytes("input.txt")
	first, err := FirstStar(f)

	if err != nil {
		fmt.Printf(`An error occured... somehow: %v`, err)
		return
	}
	fmt.Printf(`First Star Result: %v%v`, first, "\r\n")

	second, err := SecondStar(f)
	if err != nil {
		fmt.Printf(`An error occured... somehow: %v`, err)
		return
	}
	fmt.Printf(`Second Star Result: %v%v`, second, "\r\n")
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
	var res int
	numStore := []int{}
	numMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := range f {
		three := numWordBuilder(f, i, 3)
		four := numWordBuilder(f, i, 4)
		five := numWordBuilder(f, i, 5)

		switch f[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			numStore = append(numStore, int(f[i]-48))
		case 10:
			res += numStore[0]*10 + numStore[len(numStore)-1]
			numStore = []int{}
		default:
			if v, ok := numMap[three]; ok {
				numStore = append(numStore, v)
			}
			if v, ok := numMap[four]; ok {
				numStore = append(numStore, v)
			}
			if v, ok := numMap[five]; ok {
				numStore = append(numStore, v)
			}
		}

	}
	return res, nil
}

func numWordBuilder(f []byte, idx int, l int) string {
	var numWord string
	if remaining := len(f) - idx; remaining >= l {
		numWord = string(f[idx : idx+l])
	}
	return numWord
}