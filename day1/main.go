package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	var nums []int

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			nums = append(nums, i)
		}
	}

	for _, x := range nums {
		for _, y := range nums {
			if x+y == 2020 {
				fmt.Printf("%d + %d = 2020   ( %d * %d is %d )\n", x, y, x, y, x*y)
			}
		}
	}
}
