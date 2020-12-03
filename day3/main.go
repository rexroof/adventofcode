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
}
