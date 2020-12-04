package main

import (
	"bufio"
	"day4/passport"
	"fmt"
	"os"
	"strings"
)

// this function could be built into passport?
func initPeon() map[string]string {

	// I'm sure I'll think of a better way!
	return map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}

}

func checkPeon(peon map[string]string) bool {
	pass := passport.New(peon["byr"], peon["iyr"], peon["eyr"],
		peon["hgt"], peon["hcl"], peon["ecl"], peon["pid"], peon["cid"])
	return pass.CheckValidity()
}

func main() {
	// read file
	inputFile := "input.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	valid, invalid := 0, 0

	peon := initPeon()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// if this is a blank line, assume we just finished an object
		if len(text) > 0 {
			for _, chunk := range strings.Split(text, " ") {
				fmt.Printf("chunk: %s\n", chunk)
				bitty := strings.Split(chunk, ":")
				fmt.Printf("setting peon[%s] = %s\n", bitty[0], bitty[1])
				peon[bitty[0]] = bitty[1]
			}
		} else {
			if checkPeon(peon) {
				fmt.Println("this peon is valid")
				valid++
			} else {
				fmt.Println("this peon is not valid")
				invalid++
			}
			peon = initPeon()
		}

	}
	// check final peon
	if checkPeon(peon) {
		fmt.Println("this peon is valid")
		valid++
	} else {
		fmt.Println("this peon is not valid")
		invalid++
	}

	fmt.Printf("valid: %d invalid: %d\n", valid, invalid)

}
