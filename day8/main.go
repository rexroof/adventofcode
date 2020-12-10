package main

import (
	"bufio"
	"day8/instruction"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// acc , increase accumulator
// jmp , jump
// nop , nadda

/*
read all instructions into an array
track how many times each instruction has been run
if it's > 0, stop running instructions
*/

func main() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer file.Close()

	var instructionsList []instruction.Instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// could potentially use regex... this is simpler for now
		words := strings.Split(text, " ")
		operation := words[0]
		argument, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		instructionsList = append(instructionsList, instruction.New(operation, argument))
	}
	// acc , increase accumulator

	accumulator := 0
	index := 0
process:
	for {
		current := instructionsList[index]
		//fmt.Printf("acc %d idx %d op %s arg %d exec %d \n",
		//   accumulator, index, current.Operation, current.Argument, current.Executed)

		if current.Executed > 0 {
			break process
		} else {
			instructionsList[index] = current.Exec()
		}

		// acc jmp nop
		if current.Operation == "acc" {
			accumulator += current.Argument
			index++
		} else if current.Operation == "jmp" {
			index += current.Argument
		} else if current.Operation == "nop" {
			index++
		} else {
			fmt.Printf("WTF IS %s\n", current.Operation)
			break process
		}

	}

	fmt.Printf("all done, accumulator is %d\n", accumulator)

}
