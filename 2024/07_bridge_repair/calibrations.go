package main

import (
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
		if slices.Contains(results[len(results)/2:], expected) {
			sum += expected
		}
	}

	return sum
}

func calcAllPossibilities(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	// results is a flat binary tree of multiplication/adding all nums
	// left is adding, right is multi
	results := []int{nums[0]}
	for i, a, b := 1, 0, 0; i < len(nums); i, a, b = i+1, a+math.Pow2(i-1), b+math.Pow2(i) {
		s := a
		for s <= b {
			results = append(results, results[s]+nums[i], results[s]*nums[i])
			s++
		}
	}
	return results
}
