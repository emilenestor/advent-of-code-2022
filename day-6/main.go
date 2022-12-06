package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filePath := "day-6/input_test.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func distinct(s string) string {
	letters := make(map[rune]bool)
	res := ""
	for _, letter := range s {
		exists := letters[letter]
		if exists {
			continue
		}
		res += string(letter)
		letters[letter] = true
	}
	return res
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}

	num := 0
	length := 4
	for i := length - 1; i < len(line); i++ {
		checkLine := line[i-(length-1) : i+1]
		distinctLine := distinct(checkLine)
		if len(distinctLine) > length-1 {
			return i + 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return num
}

func secondStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}

	num := 0
	length := 14
	for i := length - 1; i < len(line); i++ {
		checkLine := line[i-(length-1) : i+1]
		distinctLine := distinct(checkLine)
		if len(distinctLine) > length-1 {
			return i + 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return num
}
