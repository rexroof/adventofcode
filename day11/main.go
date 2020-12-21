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

// return number of occupied seats
func countOccupied(seats [][]rune) int {
	_count := 0
	for _, row := range seats {
		for _, s := range row {
			if s == '#' {
				_count++
			}
		}
	}
	return _count
}

// returns a seat if it exists.  returns X if doesn't exist
func grabSeat(seats [][]rune, x int, y int) rune {
	_s := 'X'

	if x >= 0 && y >= 0 {
		if x < len(seats) && y < len(seats[x]) {
			_s = seats[x][y]
		}
	}
	return _s
}

// scan whole array checking and applying rule1
// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
func RuleOne(seats [][]rune) ([][]rune, bool) {
	_copyseats := [][]rune{}
	_changed := false

	for x, row := range seats {
		_copyrow := []rune{}
		for y, s := range row {
			if s == 'L' {
				// fmt.Printf("%d %d %c\n", x, y, s)
				clear := true

				// above
				if grabSeat(seats, x-1, y) == '#' {
					clear = false
				}

				// above left
				if grabSeat(seats, x-1, y-1) == '#' {
					clear = false
				}

				// above right
				if grabSeat(seats, x-1, y+1) == '#' {
					clear = false
				}

				// below
				if grabSeat(seats, x+1, y) == '#' {
					clear = false
				}

				// below left
				if grabSeat(seats, x+1, y-1) == '#' {
					clear = false
				}

				// below right
				if grabSeat(seats, x+1, y+1) == '#' {
					clear = false
				}

				// left
				if grabSeat(seats, x, y-1) == '#' {
					clear = false
				}

				// right
				if grabSeat(seats, x, y+1) == '#' {
					clear = false
				}

				if clear == true {
					_copyrow = append(_copyrow, '#')
					_changed = true
				} else {
					_copyrow = append(_copyrow, s)
				}
			} else {
				_copyrow = append(_copyrow, s)
			}
		}
		_copyseats = append(_copyseats, _copyrow)
	}
	return _copyseats, _changed
}

// scan whole array checking and applying rule2
// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
func RuleTwo(seats [][]rune) ([][]rune, bool) {
	_copyseats := [][]rune{}
	_changed := false

	for x, row := range seats {
		_copyrow := []rune{}

		for y, s := range row {
			if s == '#' {
				// fmt.Printf("%d %d %c\n", x, y, s)
				neighbors := 0

				// above
				if grabSeat(seats, x-1, y) == '#' {
					neighbors++
				}

				// above left
				if grabSeat(seats, x-1, y-1) == '#' {
					neighbors++
				}

				// above right
				if grabSeat(seats, x-1, y+1) == '#' {
					neighbors++
				}

				// below
				if grabSeat(seats, x+1, y) == '#' {
					neighbors++
				}

				// below left
				if grabSeat(seats, x+1, y-1) == '#' {
					neighbors++
				}

				// below right
				if grabSeat(seats, x+1, y+1) == '#' {
					neighbors++
				}

				// left
				if grabSeat(seats, x, y-1) == '#' {
					neighbors++
				}

				// right
				if grabSeat(seats, x, y+1) == '#' {
					neighbors++
				}

				if neighbors > 3 {
					_copyrow = append(_copyrow, 'L')
					// fmt.Printf("%d %d - changed to %c\n", x, y, s)
					_changed = true
				} else {
					_copyrow = append(_copyrow, s)
					// fmt.Printf("%d %d - kept with %c (%d) \n", x, y, s, neighbors)
				}

			} else {
				_copyrow = append(_copyrow, s)
			}
		}

		_copyseats = append(_copyseats, _copyrow)
	}
	return _copyseats, _changed
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

	fmt.Println("first few rows:")
	PrintRow(seats[0])
	PrintRow(seats[1])
	PrintRow(seats[2])
	PrintRow(seats[3])
	PrintRow(seats[4])

	// apply each rule
	seats, _ = RuleOne(seats)
	seats, _ = RuleTwo(seats)

	// if we are still moving
	stillShifting := true
	// how many times we applied the rules
	rulesCount := 2

	for stillShifting == true {
		one, two := false, false

		seats, one = RuleOne(seats)
		rulesCount++
		seats, two = RuleTwo(seats)
		rulesCount++

		if one == false && two == false {
			stillShifting = false
		}

	}

	fmt.Printf("after %d runs of the rules: \n", rulesCount)
	PrintRow(seats[0])
	PrintRow(seats[1])
	PrintRow(seats[2])
	PrintRow(seats[3])
	PrintRow(seats[4])

	fmt.Printf("counted %d seats occupied\n", countOccupied(seats))

}
