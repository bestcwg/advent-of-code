package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solvePart1() int {
	values1, values2 := loadInput()
	slices.Sort(values1)
	slices.Sort(values2)

	acc := 0
	for i := range len(values1) {
		if values1[i] > values2[i] {
			acc += values1[i] - values2[i]
		} else {
			acc += values2[i] - values1[i]
		}
	}

	return acc
}

func loadInput() ([]int, []int) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	list := strings.Split(string(data), "\n")

	list1 := make([]int, len(list))
	list2 := make([]int, len(list))

	for i := 0; i < len(list); i++ {
		split := strings.Split(list[i], "   ")
		list1[i], _ = strconv.Atoi(split[0])
		list2[i], _ = strconv.Atoi(split[1])
	}

	return list1, list2
}
