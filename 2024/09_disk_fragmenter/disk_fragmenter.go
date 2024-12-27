package main

import (
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	println("pt1 ", pt1(string(input)))
}

func prepare(in string) []int {
	var out []int
	for i, ch := range in {
		n := int(ch - 48)
		if i%2 == 0 {
			out = append(out, slices.Repeat([]int{i / 2}, n)...)
		} else {
			out = append(out, slices.Repeat([]int{-1}, n)...)
		}
	}

	return out
}

func pt1(in string) int {
	disk := prepare(in)

	sum := 0
	for i, j := 0, len(disk)-1; i <= j; {
		switch {
		case disk[i] > -1:
			sum += i * disk[i]
			i++
		case disk[j] > -1:
			sum += i * disk[j]
			j--
			i++
		default:
			j--
		}
	}
	return sum
}
