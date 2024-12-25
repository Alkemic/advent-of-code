package main

import (
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), "\n\n")
	orderingPart := parts[0]
	updatesPart := parts[1]

	orderings := map[string][]string{}
	for _, update := range strings.Split(orderingPart, "\n") {
		split := strings.Split(update, "|")
		orderings[split[0]] = append(orderings[split[0]], split[1])
	}

	var updates [][]string
	for _, part := range strings.Split(updatesPart, "\n") {
		updates = append(updates, strings.Split(part, ","))
	}

	sum := 0
	for _, update := range updates {
		fiexed := false
		for ok, idx1, idx2 := checkOrderingRule(update, orderings); !ok; ok, idx1, idx2 = checkOrderingRule(update, orderings) {
			fiexed = true
			// swap places
			update[idx1], update[idx2] = update[idx2], update[idx1]
		}

		if fiexed {
			mid, _ := strconv.Atoi(update[len(update)/2])
			sum += mid
		}
	}

	println("sum", sum)
}

func checkOrderingRule(updatePages []string, orderings map[string][]string) (bool, int, int) {
	// for each update page, check for rules, and if ordering is correct
	for i := 0; i < len(updatePages); i++ {
		page := updatePages[i]
		ordering := orderings[page]
		for _, after := range ordering {
			if slices.Contains(updatePages[:i], after) {
				return false, i, slices.Index(updatePages[:i], after)
			}
		}
	}
	return true, -1, -1
}
