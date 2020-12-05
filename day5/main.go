package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*

// 128 rows, F - front half, B - back half, repeating

128
F
1-64            65-128
F
1-32 33-64
F
1-16 17-32
F
1-8 9-16
F
1-4 5-8
F
1-2 3-4
F B F B
1 2 3 4


*/

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

func calcSeatID(boardingPass string) int64 {
	// rowNum := rowToNum(boardingPass[0:7])
	// posNum := rowToNum(boardingPass[7:10])
	// multiply the row by 8, then add the column
	// return ((rowNum * 8) + posNum)
	return ((rowToNum(boardingPass[0:7]) * 8) + rowToNum(boardingPass[7:10]))
}

func main() {
	a := "FBFBBFFRLR"
	fmt.Printf("%s split is %s and %s\n", a, a[0:7], a[7:10])
	fmt.Printf("%s SeatID is %d\n", a, calcSeatID(a))

	a = "BFFFBBFRRR"
	fmt.Printf("%s split is %s and %s\n", a, a[0:7], a[7:10])
	fmt.Printf("%s SeatID is %d\n", a, calcSeatID(a))
	a = "FFFBBBFRRR"
	fmt.Printf("%s split is %s and %s\n", a, a[0:7], a[7:10])
	fmt.Printf("%s SeatID is %d\n", a, calcSeatID(a))
	a = "BBFFBBFRLL"
	fmt.Printf("%s split is %s and %s\n", a, a[0:7], a[7:10])
	fmt.Printf("%s SeatID is %d\n", a, calcSeatID(a))

	fmt.Printf("BBBBBBB is %d\n", rowToNum("BBBBBBB"))
	fmt.Printf("FFFFFFF is %d\n", rowToNum("FFFFFFF"))
	fmt.Printf("FBFBBFF is %d\n", rowToNum("FBFBBFF"))
	fmt.Printf("RRR is %d\n", rowToNum("RRR"))
	fmt.Printf("RLR is %d\n", rowToNum("RLR"))
	fmt.Printf("LLL is %d\n", rowToNum("LLL"))

}
