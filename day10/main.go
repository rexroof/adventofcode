package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	var joltyAdapters []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		joltyAdapters = append(joltyAdapters, value)
	}

	sort.Ints(joltyAdapters)

	// my device is always 3 higher than the highest adapter, so we add it to the list
	joltyAdapters = append(joltyAdapters, joltyAdapters[len(joltyAdapters)-1]+3)

	// keeping this count, might be useful later
	diffs := []int{}
	for x, jolts := range joltyAdapters {
		if x == 0 {
			// use 1 to represent the jolt different to the 0 rated outlet
			diffs = append(diffs, 1)
		} else {
			diffs = append(diffs, jolts-joltyAdapters[x-1])
		}
	}

	// count one-jolt diffs and three-jolt diffs
	threeJolts, oneJolts, twoJolts := 0, 0, 0

	for _, d := range diffs {
		if d == 1 {
			oneJolts++
		} else if d == 3 {
			threeJolts++
		} else if d == 2 {
			twoJolts++
		} else if d > 3 {
			fmt.Println("btw, have a d of %d\n", d)
		}
	}

	fmt.Printf("one jolts times three jolts is  %d x %d = %d\n", oneJolts, threeJolts, oneJolts*threeJolts)
	fmt.Printf("two jolts is %d\n", twoJolts)
	fmt.Printf("is part2 just oneJolts ** 3  = %f (it is not)\n", math.Pow(float64(oneJolts), float64(3)))

	// part 2.   took hint from reddit.
	//  - breaking list of jolts up into group separated by 1 or 3
	var joltGroups [][]int
	var _tmpGroup = []int{}

	// lets put the outlet (0) on the front this array
	joltyAdapters = append([]int{0}, joltyAdapters...)

	for idx, val := range joltyAdapters {
		nextJump := -1
		if idx+1 < len(joltyAdapters) {
			nextJump = (joltyAdapters[idx+1] - val)
		}
		_tmpGroup = append(_tmpGroup, val)
		if nextJump == 3 {
			joltGroups = append(joltGroups, _tmpGroup)
			_tmpGroup = []int{}
		}
	}

	independents := 0
	triplets := 0

	for _, val := range joltGroups {
		l := len(val)

		if (l - 2) == 3 {
			triplets++
		} else if (l - 2) > 0 {
			independents += (l - 2)
		}
		// fmt.Println(idx, val, l)
	}

	ipow := math.Pow(2, float64(independents))
	tpow := math.Pow(7, float64(triplets))

	fmt.Printf("%d independents, 2^%d is %f\n", independents, independents, ipow)
	fmt.Printf("%d triplets, 7^%d is %f\n", triplets, triplets, tpow)
	fmt.Printf("whole answer might be %f\n", ipow*tpow)
}
