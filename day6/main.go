package main

import (
	"bufio"
	"fmt"
	"os"
)

func countAndPurge(m map[rune]int) (int, map[rune]int) {

	toats := 0
	for k := range m {
		toats += m[k]
		delete(m, k)
	}

	return toats, m

}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	grand_total := 0
	peons := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) > 0 {
			// fmt.Println(text)
			for _, char := range text {
				peons[char] = 1
				// fmt.Println(peons)
			}
		} else {
			// found blank line.  calc totals and reset map
			var toats int
			toats, peons = countAndPurge(peons)
			//fmt.Printf("our total was %d\n", toats)
			grand_total += toats
		}
	}
	var toats int
	toats, peons = countAndPurge(peons)
	//fmt.Printf("our total was %d\n", toats)
	grand_total += toats
	fmt.Printf("grand total: %d\n", grand_total)
}
