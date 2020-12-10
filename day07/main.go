package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// count how many bags are in a specific color bag
func countBagsInside(bags map[string]map[string]int, bagToCheck map[string]int) int {
	var total_bags int = 0

	// if this bag contains no other bags, we just return 1
	if count, exists := bagToCheck["none"]; exists {
		if count == 0 {
			total_bags = 0
		}
	} else {
		for color, count := range bagToCheck {
			check := countBagsInside(bags, bags[color])

			// add the number of bags.
			total_bags += count

			// plus the bags inside of those
			total_bags += (count * check)

		}
	}

	return total_bags
}

// checks to see if seekingColor is inside bagToCheck, uses bags for recursive search
func bagIsInside(bags map[string]map[string]int, bagToCheck map[string]int, seekingColor string) bool {
	_found := false
	// bagIsInside( bags, bags["vibrant black"], "shiny gold" )

	if _, exists := bagToCheck[seekingColor]; exists {
		_found = true
	} else {
		for oneBag, _ := range bagToCheck {
			// fmt.Printf("recursive call to bagIsInside(%s)\n", oneBag)
			if bagIsInside(bags, bags[oneBag], seekingColor) {
				_found = true
			}
		}
	}

	return _found
}

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	var bagsDef = map[string]map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// could potentially use regex... this is simpler for now
		words := strings.Split(text, " ")

		// example
		// posh violet bags contain 2 dark violet bags, 4 striped olive bags, 1 pale silver bag.

		bag_color := strings.Join(words[0:2], " ")
		// remove bag color from array.
		words = words[2:]
		// remove "bag(s) contain(s)"
		words = words[2:]

		bagsDef[bag_color] = make(map[string]int)

		if strings.Contains(text, "contain no other bags") {
			bagsDef[bag_color]["none"] = 0
		} else {
			for len(words) > 2 {
				if bag_count, err := strconv.Atoi(words[0]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else if bag_count > 0 {
					sub_color := strings.Join(words[1:3], " ")
					bagsDef[bag_color][sub_color] = bag_count
					// fmt.Printf("bagsDef[%s][%s] = %d\n", bag_color, sub_color, bag_count)

					// remove count, colors, then the word "bags,|."
					words = words[4:]
				}
			}
		}

	}

	total_found_in := 0

	for key, _ := range bagsDef {
		// fmt.Printf("checking %s:\n", key)

		if bagIsInside(bagsDef, bagsDef[key], "shiny gold") {
			// fmt.Printf("FOUND IN %s\n", key)
			total_found_in++
		} else {
			// fmt.Printf("NOT FOUND IN %s\n", key)
		}

	}

	fmt.Printf("total found shiny gold in: %d\n", total_found_in)
	fmt.Printf("shiny gold contains a total of %d bags\n", countBagsInside(bagsDef, bagsDef["shiny gold"]))
}
