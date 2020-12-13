package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// return array with all elements except the index passed
func filterSlice(input []int, index int) []int {
	var truck []int

	// if this is the first element...
	if index == 0 {
		truck = make([]int, len(input)-1)
		copy(truck, input[1:len(input)])
		// if this is the last element...
	} else if index == (len(input) - 1) {
		truck = make([]int, len(input)-1)
		copy(truck, input[0:len(input)-1])
	} else {
		truck = make([]int, len(input[0:index]))
		copy(truck, input[0:index])
		truck = append(truck, input[index+1:len(input)]...)
	}

	return truck
}

func passesSumTest(list []int, value int) bool {
	passed := false
	var checking []int
	copy(checking, list)

	// for each element of our list
	for x, a := range list {
		// check every other element of this list!
		for _, b := range filterSlice(list, x) {
			fmt.Printf("checking if %d + %d = %d\n", a, b, value)
			if a+b == value {
				passed = true
			}
		}
	}

	return passed
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	var xmas []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		xmas = append(xmas, value)
	}

	groupSize := 25
	for x := groupSize; x < len(xmas); x++ {
		fmt.Printf("checking %d\n", xmas[x])
		if passesSumTest(xmas[(x-groupSize):x], xmas[x]) {
			// fmt.Println(xmas[0:10])
		} else {
			// fmt.Println(xmas[0:10])
			// fmt.Println(xmas[(x - groupSize):x])
			fmt.Printf("failed: %d\n", xmas[x])
		}
	}

}
