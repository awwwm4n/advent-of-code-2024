package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	part1()
	part2()

}

func getInputs(list1 *[]int, list2 *[]int) {
	file, err := os.Open("../inputs/day1")
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		panic("Error reading input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Fields(line)
		if len(row) < 2 {
			continue
		}

		val1, err1 := strconv.Atoi(row[0])
		val2, err2 := strconv.Atoi(row[1])

		if err1 != nil || err2 != nil {
			fmt.Printf("Invalid input: %s\n", line)
			panic("Error reading input")
		}

		*list1 = append(*list1, val1)
		*list2 = append(*list2, val2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v", err)
		panic("Error reading input")
	}
}

func part1() {
	var list1 []int
	var list2 []int

	getInputs(&list1, &list2)

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i := range list1 {
		totalDistance += int(math.Abs((float64(list1[i] - list2[i]))))
	}

	fmt.Printf("\nPart1 ans: %v", totalDistance)

}

func part2() {
	var list1 []int
	var list2 []int

	getInputs(&list1, &list2)

	sort.Ints(list1)
	sort.Ints(list2)

	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, val := range list1 {
		m1[val]++
	}

	for _, val := range list2 {
		m2[val]++
	}

	similarityScore := 0

	for k, v := range m1 {
		if m2Val, isPresent := m2[k]; isPresent {
			similarityScore += k * v * m2Val
		}
	}

	fmt.Printf("\nPart2 ans: %v", similarityScore)
}
