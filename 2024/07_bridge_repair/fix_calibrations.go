package main

import (
	math2 "math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Alkemic/aoc/math"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	println("calibrations sum:", calibrations(string(input)))
}

func concatNumbers[T int | uint | uint8 | uint16 | uint32 | uint64](a, b T) T {
	order := int(math2.Log10(float64(b))) + 1
	a *= math.Pow10(T(order))
	a += b
	return a
}

func calibrations(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ": ")

		var nums []int
		for _, n := range strings.Split(parts[1], " ") {
			r, _ := strconv.Atoi(n)
			nums = append(nums, r)
		}

		expected, _ := strconv.Atoi(parts[0])
		results := calcAllPossibilities(nums)
		// check leaves only for matching results
		if slices.Contains(results[len(results)/3:], expected) {
			//fmt.Println(expected)
			sum += expected
		}
	}

	return sum
}

type ops struct {
	operations []func(a, b int) int
}

var x = ops{
	operations: []func(a, b int) int{
		func(a, b int) int {
			return a + b
		},
		func(a, b int) int {
			return a * b
		},
		func(a, b int) int {
			return concatNumbers(a, b)
		},
	},
}

func calcAllPossibilities(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	// results is a flat trinary tree of multiplication/adding all nums
	// left is adding, right is multi
	results := []int{nums[0]}
	for i, a, b := 1, 0, 0; i < len(nums); i, a, b = i+1, a+math.Pow3(i-1), b+math.Pow3(i) {
		s := a
		for s <= b {
			results = append(results, results[s]+nums[i], results[s]*nums[i], concatNumbers(results[s], nums[i]))
			s++
		}
	}
	return results
}
