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

	// quick sanity check.. do any nop operations contain an argument that would get us to the end?
	/* wasnt' this easy
	for x, ins := range instructionsList {
		if ins.Operation == "nop" {
			if (x + ins.Argument) == len(instructionsList) {
				fmt.Println("found one!")
				fmt.Println(x)
				fmt.Println(ins)
			}
		}
	}
	fmt.Println("exited")
	os.Exit(0)
	*/

	accumulator := 0
	index := 0
	var stack [][]int
process:
	for {
		// if our index is beyond our instructions, we're done
		if index >= len(instructionsList) {
			break process
		}
		current := instructionsList[index]
		fmt.Printf("acc %d idx %d op %s arg %d exec %d \n",
			accumulator, index, current.Operation, current.Argument, current.Executed)
		stack = append(stack, []int{accumulator, index})
		fmt.Println(stack[len(stack)-1])

		if current.Executed > 0 {
			fmt.Printf("hit repeat.  accumulator is %d\n", accumulator)
			// if we've already executed this, rewind using the stack

			rewind := stack[len(stack)-2]
			fmt.Printf("rewound stack contains: acc %d idx %d op %s arg %d exec %d \n",
				rewind[0], rewind[1], instructionsList[rewind[1]].Operation,
				instructionsList[rewind[1]].Argument, instructionsList[rewind[1]].Executed)

			//replace rewind location with replaced jmp/nop
			if instructionsList[rewind[1]].Operation == "jmp" {
				_arg := instructionsList[rewind[1]].Argument
				instructionsList[rewind[1]] = instruction.New("nop", _arg)
				index = rewind[1]
				continue process
			} else if instructionsList[rewind[1]].Operation == "nop" {
				_arg := instructionsList[rewind[1]].Argument
				instructionsList[rewind[1]] = instruction.New("jmp", _arg)
				index = rewind[1]
				continue process
			} else {
				fmt.Printf("operation is %s, not something we can fix", instructionsList[rewind[1]].Operation)
				break process
			}
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
	fmt.Printf("program exiting, accumulator is %d\n", accumulator)
}
