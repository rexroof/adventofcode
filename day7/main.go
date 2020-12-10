package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// checks to see if seekingColor is inside bagToCheck, uses bags for recursive search
func bagIsInside(bags map[string]map[string]int, bagToCheck map[string]int, seekingColor string) bool {
	_found := false
	// bagIsInside( bags, bags["vibrant black"], "shiny gold" )

	if _, exists := bagToCheck[seekingColor]; exists {
		_found = true
	} else {
		for oneBag, _ := range bagToCheck {
			fmt.Printf("recursive call to bagIsInside(%s)\n", oneBag)
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
				} else {
					sub_color := strings.Join(words[1:3], " ")
					bagsDef[bag_color][sub_color] = bag_count

					// remove count, colors, then the word "bags,|."
					words = words[4:]
				}
			}
		}

	}

	total_found_in := 0

	for key, _ := range bagsDef {
		fmt.Printf("checking %s:\n", key)

		if bagIsInside(bagsDef, bagsDef[key], "shiny gold") {
			fmt.Printf("FOUND IN %s\n", key)
			total_found_in++
		} else {
			fmt.Printf("NOT FOUND IN %s\n", key)
		}

	}

	fmt.Printf("total found shiny gold in: %d\n", total_found_in)

}
