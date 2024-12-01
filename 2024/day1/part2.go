package main

import (
	"slices"
)

func solvePart2() int {
	values1, values2 := loadInput()
	slices.Sort(values1)
	slices.Sort(values2)

	acc := 0
	for i := range len(values1) {
		count := 0
		for j := range len(values2) {
			if values2[j] == values1[i] {
				count += 1
			}
		}

		if count != 0 {
			acc += values1[i] * count
		}
	}

	return acc
}
