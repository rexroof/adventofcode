package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	r, err := regexp.Compile(`([\d]+)-([\d]+) ([\w]): ([\w]+)`)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	valid := 0
	invalid := 0

	for _, line := range lines {
		// fmt.Println(line)
		results := r.FindSubmatch([]byte(line))

		/*  [ 0 "9-11 b: bkbltdvbtwbbtsb"
		    1 "9"
		    2 "11"
		    3 "b"
		    4 "bkbltdvbtwbbtsb"] */

		letMin, err := strconv.Atoi(string(results[1]))
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		letMax, err := strconv.Atoi(string(results[2]))
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		let := string(results[3])
		pword := string(results[4])

		reglet := regexp.MustCompile(let)
		matchlet := reglet.FindAllStringIndex(pword, -1)

		if len(matchlet) >= letMin && len(matchlet) <= letMax {
			valid++
		} else {
			invalid++
			// fmt.Printf("found %d matches\n", len(matchlet))
			//fmt.Println(pword)
		}

	}
	fmt.Printf("valid: %d  invalid: %d\n", valid, invalid)

}
