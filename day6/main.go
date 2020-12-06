package main

import (
	"bufio"
	"fmt"
	"os"
)

func peonPurge(m map[rune]int) map[rune]int {
	for k := range m {
		delete(m, k)
	}
	return m
}

func countKeys(m map[rune]int) int {
	toats := 0
	for range m {
		toats++
	}
	return toats
}

func countOnlyAll(count int, m map[rune]int) int {
	toats := 0
	for k := range m {
		// fmt.Printf(" m[%c] = %d (count is %d) \n", k, m[k], count)
		if m[k] == count {
			toats++
		}
	}
	return toats
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	any_answer_total := 0
	all_answer_total := 0
	peons := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	numPeeps := 0
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) > 0 {
			// fmt.Println(text)
			numPeeps++
			for _, char := range text {
				peons[char]++
				// fmt.Println(peons)
			}
		} else {
			// found blank line.  calc totals and reset map
			any_answer_total += countKeys(peons)
			all_answer_total += countOnlyAll(numPeeps, peons)

			numPeeps = 0
			peons = peonPurge(peons)
		}
	}
	// count last object
	any_answer_total += countKeys(peons)
	all_answer_total += countOnlyAll(numPeeps, peons)

	fmt.Printf("any answer total: %d\n", any_answer_total)
	fmt.Printf("all answer total: %d\n", all_answer_total)
}
