package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile := "inputs/2.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Could not read file %s - ERR: %v\n", inputFile, err)
		os.Exit(1)
	}

	safeCounter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
		if isSafe(report) {
			safeCounter++
		}
	}
	fmt.Printf("Safe Counter: %d reports!\n", safeCounter)
}

func isGradual(report []string) bool {
	if len(report) == 1 {
		return true
	} else {
		r1, _ := strconv.Atoi(report[1])
		r0, _ := strconv.Atoi(report[0])

		var signal int
		if r1-r0 > 0 {
			signal = 1
		} else if r1-r0 < 0 {
			signal = -1
		} else {
			return false
		}

		for i, _ := range report {
			if i == 0 || i == 1 {
				continue
			}

			r1, _ = strconv.Atoi(report[i])
			r0, _ = strconv.Atoi(report[i-1])
			if (r1-r0)*signal < 0 {
				return false
			}
		}
	}
	return true
}

func isDiffOk(report []string) bool {
	if len(report) == 1 {
		return true
	} else {
		for i, _ := range report {
			if i == 0 {
				continue
			}
			r1, _ := strconv.Atoi(report[i])
			r0, _ := strconv.Atoi(report[i-1])
			diff := r1 - r0
			if diff < 0 {
				diff = diff * -1
			}
			if diff == 0 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func isSafe(report []string) bool {
	if isGradual(report) && isDiffOk(report) {
		return true
	} else {
		// This is probably not the most efficient way O(n^2)
		for i, _ := range report {
			newReport := removeLevel(report, i)
			if isGradual(newReport) && isDiffOk(newReport) {
				return true
			}
		}
		return false
	}
}

func removeLevel(report []string, levelIndex int) []string {
	return slices.Concat(report[0:levelIndex], report[levelIndex+1:])
}
