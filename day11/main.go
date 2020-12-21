package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
. floor
L empty
# occupied

If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/

// print one row
func PrintRow(row []rune) {
	fmt.Println(string(row))
}

// scan whole array checking and applying rule1
// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
func RuleOne(seats [][]rune) [][]rune {
	for x, row := range seats {
		for y, s := range row {
			if s == 'L' {
				fmt.Printf("%d %d %c\n", x, y, s)
				clear := true

				// above
				if x-1 >= 0 {
					abv := seats[x-1][y]
					if abv == '#' {
						clear = false
					}
				}

				// below
				if x+1 < len(seats) {
					blw := seats[x+1][y]
					if blw == '#' {
						clear = false
					}
				}

				// left
				if y-1 >= 0 {
					lft := seats[x][y-1]
					if lft == '#' {
						clear = false
					}
				}

				// right
				if y+1 < len(seats[x]) {
					rgt := seats[x][y+1]
					if rgt == '#' {
						clear = false
					}
				}

				if clear == true {
					seats[x][y] = '#'
				}

			}
		}
	}
	return seats
}

// scan whole array checking and applying rule1
// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
func RuleTwo(seats [][]rune) [][]rune {
	return seats
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	seats := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		oneRow := []rune{}
		for _, char := range text {
			oneRow = append(oneRow, char)
		}

		seats = append(seats, oneRow)
	}

	seats = RuleOne(seats)

	PrintRow(seats[5])
	PrintRow(seats[6])
	PrintRow(seats[7])
	PrintRow(seats[8])

}
