package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Starting at the top-left corner of your map and following
  a slope of right 3 and down 1, how many trees would you encounter?
*/

func countTrees(slope [][]bool, offset_x int, offset_y int) int {
	// starting spot
	x, y := 0, 0
	tree_count := 0

	// this isn't great for x > 1, but the break fixes it
	for i := 1; i < len(slope); i++ {
		myx := x + offset_x*i
		myy := y + offset_y*i

		if myx > len(slope) {
			break
		}

		// since our pattern of this line repeats,
		//   shift our index to the left by the length of our array
		for myy >= len(slope[myx]) {
			myy -= len(slope[myx])
		}

		if !(slope[myx][myy]) {
			tree_count++
		}
	}
	return tree_count
}

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	// load our file into an array of arrays
	var theSlope [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var oneRow []bool
		for _, c := range scanner.Text() {
			if c == '.' {
				// clear spot is true
				oneRow = append(oneRow, true)
			} else if c == '#' {
				// tree is false
				oneRow = append(oneRow, false)
			} else {
				fmt.Printf("wtf is this character? %c", c)
			}
		}
		theSlope = append(theSlope, oneRow)
	}

	checking := [5][2]int{
		{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}

	sum := 1
	for _, a := range checking {
		trees := countTrees(theSlope, a[0], a[1])
		fmt.Printf("slope: %d %d trees: %d\n", a[0], a[1], trees)
		sum *= trees
	}
	fmt.Printf("sum: %d\n", sum)

}
