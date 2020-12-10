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

// process a set of instructions.
//   return false and value if it repeats,
//   return true and value if it reaches end
func processInstructions(inst []instruction.Instruction) (bool, int) {
	_acc := 0
	index := 0
	for {
		// if our index is beyond our instructions, we're done
		if index >= len(inst) {
			return true, _acc
		}
		current := inst[index]

		if current.Executed > 0 {
			// this was already run, return false
			return false, _acc
		} else {
			inst[index] = current.Exec()
		}

		if current.Operation == "acc" {
			_acc += current.Argument
			index++
		} else if current.Operation == "jmp" {
			index += current.Argument
		} else if current.Operation == "nop" {
			index++
		} else {
			fmt.Printf("WTF IS %s\n", current.Operation)
			return false, -1
		}
	}
}

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

	// make copy and run test on original list copy
	_copy := make([]instruction.Instruction, len(instructionsList))
	for i := range instructionsList {
		_copy[i] = instructionsList[i]
	}

	pass, count := processInstructions(_copy)
	fmt.Printf("original results:  %t, %d\n", pass, count)

	for x, inst := range instructionsList {
		if inst.Operation == "nop" {

			// make a copy of our instructions
			_copy := make([]instruction.Instruction, len(instructionsList))
			for i := range instructionsList {
				_copy[i] = instructionsList[i]
			}

			_copy[x] = instruction.New("jmp", inst.Argument)
			if pass, count := processInstructions(_copy); pass {
				fmt.Printf("omg, we got a pass with accumulator: %d\n", count)
			} else {
				// fmt.Printf("failed results: %t %d\n", pass, count)
			}

		} else if inst.Operation == "jmp" {

			// make a copy of our instructions
			_copy := make([]instruction.Instruction, len(instructionsList))
			for i := range instructionsList {
				_copy[i] = instructionsList[i]
			}

			_copy[x] = instruction.New("nop", inst.Argument)
			if pass, count := processInstructions(_copy); pass {
				fmt.Printf("omg, we got a pass with accumulator: %d\n", count)
			} else {
				// fmt.Printf("failed results: %t %d\n", pass, count)
			}
		}
	}

}
