package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	var joltyAdapters []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		joltyAdapters = append(joltyAdapters, value)
	}

	sort.Ints(joltyAdapters)

	// my device is always 3 higher than the highest adapter, so we add it to the list
	joltyAdapters = append(joltyAdapters, joltyAdapters[len(joltyAdapters)-1]+3)

	// keeping this count, might be useful later
	diffs := []int{}
	for x, jolts := range joltyAdapters {
		if x == 0 {
			// use 1 to represent the jolt different to the 0 rated outlet
			diffs = append(diffs, 1)
		} else {
			diffs = append(diffs, jolts-joltyAdapters[x-1])
		}
	}

	// count one-jolt diffs and three-jolt diffs
	threeJolts, oneJolts := 0, 0

	for _, d := range diffs {
		if d == 1 {
			oneJolts++
		} else if d == 3 {
			threeJolts++
		} else if d > 3 {
			fmt.Println("btw, have a d of %d\n", d)
		}

	}

	fmt.Printf("one jolts times three jolts is  %d x %d = %d\n", oneJolts, threeJolts, oneJolts*threeJolts)

}
