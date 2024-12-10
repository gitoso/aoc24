// https://adventofcode.com/2024/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := "inputs/1.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Could not read file %s - ERR: %v\n", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "   ")

		n1, _ := strconv.Atoi(splitted[0])
		list1 = append(list1, n1)

		n2, _ := strconv.Atoi(splitted[1])
		list2 = append(list2, n2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		dist := list1[i] - list2[i]
		if dist < 0 {
			dist = dist * -1
		}

		totalDistance += dist
	}

	fmt.Printf("Part1 - Total Distance: %d\n", totalDistance)

	totalScore := 0
	for _, n1 := range list1 {
		numTimes := 0

		for _, n2 := range list2 {
			if n1 == n2 {
				numTimes++
			}
		}
		totalScore += n1 * numTimes
	}

	fmt.Printf("Part2 - Total Score: %d\n", totalScore)
}
