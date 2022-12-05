package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type CrateStack []string

var CrateStacks []CrateStack

// IsEmpty: check if stack is empty
func (c *CrateStack) IsEmpty() bool {
	return len(*c) == 0
}

// Push a new value onto the stack
func (c *CrateStack) Push(str string) {
	*c = append(*c, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (c *CrateStack) Pop() (string, bool) {
	if c.IsEmpty() {
		return "", false
	} else {
		index := len(*c) - 1   // Get the index of the top most element.
		element := (*c)[index] // Index into the slice and obtain the element.
		*c = (*c)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	filePath := "day-5/input_test.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func getCrateStacks(scanner *bufio.Scanner) []CrateStack {
	crates := []CrateStack{}
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			for _, stack := range crates {
				reverse(stack)
			}
			return crates
		}
		stack := 0
		for i := 1; i < len(line); i += 4 {
			re := regexp.MustCompile("[A-Z]+")
			if re.MatchString(string(line[i])) {
				if firstLine {
					crates = make([]CrateStack, (len(line)-2)/3)
					firstLine = false
				}
				if crates[stack] == nil {
					crates[stack] = make([]string, 0)
				}
				crates[stack].Push(string(line[i]))
			}
			stack++
		}
		fmt.Println(line)
	}

	return crates
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func getInstructions(scanner *bufio.Scanner) [][]int {
	instructions := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		intArr := getInstructionInts(line)

		instructions = append(instructions, intArr)
	}

	return instructions
}

func getInstructionInts(instructionString string) []int {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(instructionString, -1)

	intArr := make([]int, 0)
	for _, num := range nums {
		intNum, _ := strconv.Atoi(num)
		intArr = append(intArr, intNum)
	}

	return intArr
}

func runCrateMover9000(crates *[]CrateStack, instructions *[][]int) string {
	result := ""
	for _, instruction := range *instructions {
		for i := 0; i < instruction[0]; i++ {
			temp, success := (*crates)[instruction[1]-1].Pop()

			if success {
				(*crates)[instruction[2]-1].Push(temp)
			}
		}
	}

	for _, stack := range *crates {
		if stack.IsEmpty() {
			continue
		}
		result += stack[len(stack)-1]
	}

	return result
}

func runCrateMover9001(crates *[]CrateStack, instructions *[][]int) string {
	result := ""
	for _, instruction := range *instructions {
		// Grab crates
		temp := (*crates)[instruction[1]-1][len((*crates)[instruction[1]-1])-instruction[0]:]

		// Remove crates
		(*crates)[instruction[1]-1] = (*crates)[instruction[1]-1][:len((*crates)[instruction[1]-1])-instruction[0]]

		// Append crates to another stack
		(*crates)[instruction[2]-1] = append((*crates)[instruction[2]-1], temp...)
	}

	for _, stack := range *crates {
		if stack.IsEmpty() {
			continue
		}
		result += stack[len(stack)-1]
	}

	return result
}

func firstStar(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	CrateStacks = getCrateStacks(scanner)

	instructions := getInstructions(scanner)

	result := runCrateMover9000(&CrateStacks, &instructions)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func secondStar(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	CrateStacks = getCrateStacks(scanner)

	instructions := getInstructions(scanner)

	result := runCrateMover9001(&CrateStacks, &instructions)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
