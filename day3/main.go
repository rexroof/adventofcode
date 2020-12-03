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

	/*
		for _, arr := range theSlope {
			fmt.Println(arr)
		}
	*/

	// starting spot
	x, y := 0, 0

	// slope
	xo, yo := 1, 3

	clear_count := 0
	tree_count := 0

	// for i := 1; i <= 100; i++ {
	for i := 1; i < len(theSlope); i++ {
		myx := x + xo*i
		myy := y + yo*i

		fmt.Printf("%d,%d\n", myx, myy)

		// since our pattern of this line repeats,
		//   shift our index to the left by the length of our array
		//   ( -1 to account for 0 )
		for myy >= (len(theSlope[myx]) - 1) {
			myy -= (len(theSlope[myx]) - 1)
		}

		if theSlope[myx][myy] {
			clear_count++
		} else {
			tree_count++
		}

		fmt.Printf("%d,%d %t\n", myx, myy, theSlope[myx][myy])

	}

	fmt.Printf("clear: %d tree: %d\n", clear_count, tree_count)

}
