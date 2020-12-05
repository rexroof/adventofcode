package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// this converts rows and seats to numbers
func rowToNum(row string) int64 {
	// replace F/B with 0/1
	rowBinary := strings.Replace(row, "B", "1", -1)
	rowBinary = strings.Replace(rowBinary, "F", "0", -1)
	rowBinary = strings.Replace(rowBinary, "L", "0", -1)
	rowBinary = strings.Replace(rowBinary, "R", "1", -1)

	// convert binary string to int
	if i, err := strconv.ParseInt(rowBinary, 2, 8); err != nil {
		fmt.Println(err)
		return -1
	} else {
		return i
	}
}

// this generates our seatID
func calcSeatID(boardingPass string) int64 {
	// rowNum := rowToNum(boardingPass[0:7])
	// posNum := rowToNum(boardingPass[7:10])
	// multiply the row by 8, then add the column
	// return ((rowNum * 8) + posNum)
	return ((rowToNum(boardingPass[0:7]) * 8) + rowToNum(boardingPass[7:10]))
}

func main() {
	inputFile := "input.txt"

	// open file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// to try and track missing seat, looking for min, max, sum

	// read through file, track largest seatid in biggest
	var biggest float64
	var smallest float64 = 99999999
	var toats float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seatid := calcSeatID(scanner.Text())
		// fmt.Println(seatid)

		biggest = math.Max(biggest, float64(seatid))
		toats = (toats + float64(seatid))
		smallest = math.Min(smallest, float64(seatid))
	}

	fmt.Printf("largest SeatID was %g\n", biggest)
	fmt.Printf("smallest SeatID was %g\n", smallest)
	fmt.Printf("total SeatIDs was %g\n", toats)

	// expected sum sequence up to highest
	hightotal := (biggest * (biggest + 1)) / 2
	// sum of number up to our smallest
	smalltotal := ((smallest - 1) * smallest) / 2
	expectedsum := hightotal - smalltotal

	// difference here *SHOULD* be missing seat
	fmt.Printf("expected sum: %g difference: %g\n", expectedsum, (expectedsum - toats))
}
