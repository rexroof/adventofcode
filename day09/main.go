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

// start at list[index], sum numbers to reach target,
// return true if we can match target match
func contiguousSearch(list []int, index int, target int) (bool, int, int) {
	scratch := 0
	localMin := 999999999999999
	localMax := 0

	for _, num := range list[index:len(list)] {
		scratch += num

		if num > localMax {
			localMax = num
		}
		if num < localMin {
			localMin = num
		}

		// if we hit the target, we done!
		if scratch == target {
			return true, localMin, localMax
			// if we passed the target, we done!
		} else if scratch > target {
			return false, -1, -1
		}
	}

	return false, -1, -1
}

func passesSumTest(list []int, value int) bool {
	passed := false
	var checking []int
	copy(checking, list)

	// for each element of our list
	for x, a := range list {
		// check every other element of this list!
		for _, b := range filterSlice(list, x) {
			// fmt.Printf("checking if %d + %d = %d\n", a, b, value)
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

	failed := 0

	// first find our number that doesn't have a sum set in it's group
	groupSize := 25
	for x := groupSize; x < len(xmas); x++ {
		// fmt.Printf("checking %d\n", xmas[x])
		if passesSumTest(xmas[(x-groupSize):x], xmas[x]) {
			// fmt.Println(xmas[0:10])
		} else {
			// fmt.Println(xmas[0:10])
			// fmt.Println(xmas[(x - groupSize):x])
			// fmt.Printf("failed: %d\n", xmas[x])
			failed = xmas[x]
		}
	}

	fmt.Printf("our failed number is %d, searching array\n", failed)

	for idx, val := range xmas {
		if val != failed {
			if passed, tiny, huge := contiguousSearch(xmas, idx, failed); passed {
				fmt.Printf("passed with %d + %d = %d\n", tiny, huge, tiny+huge)
			}
		}
	}

}
