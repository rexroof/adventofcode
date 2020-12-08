package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

		fmt.Printf("%s:\n", bag_color)
		if strings.Contains(text, "contain no other bags") {
			fmt.Printf("    none\n")
			bagsDef[bag_color]["none"] = 0
		} else {

			for len(words) > 2 {
				if bag_count, err := strconv.Atoi(words[0]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else {
					sub_color := strings.Join(words[1:3], " ")
					fmt.Printf("    %s(%d)\n", sub_color, bag_count)
					bagsDef[bag_color][sub_color] = bag_count

					// remove count, colors, then the word "bags,|."
					words = words[4:]
				}
			}
		}
	}

}
