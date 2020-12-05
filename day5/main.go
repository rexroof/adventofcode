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

func rowToNum(row string) int64 {
	// replace F/B with 0/1
	rowBinary := strings.Replace(row, "B", "1", -1)
	rowBinary = strings.Replace(rowBinary, "F", "0", -1)

	// convert binary string to int
	if i, err := strconv.ParseInt(rowBinary, 2, 8); err != nil {
		fmt.Println(err)
		return -1
	} else {
		return i
	}
}

func main() {

	fmt.Printf("BBBBBBB is %d\n", rowToNum("BBBBBBB"))
	fmt.Printf("FFFFFFF is %d\n", rowToNum("FFFFFFF"))
	fmt.Printf("FBFBBFF is %d\n", rowToNum("FBFBBFF"))

}
