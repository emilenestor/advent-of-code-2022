package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	filePath := "day-6/input_test.txt"

	line := getLine(filePath)

	log.Print("First Star: ", getResult(line, 4))
	log.Print("Second Star: ", getResult(line, 14))
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

func getLine(filePath string) string {
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}

func getResult(line string, length int) int {
	num := 0
	for i := length - 1; i < len(line); i++ {
		checkLine := line[i-(length-1) : i+1]
		distinctLine := distinct(checkLine)
		if len(distinctLine) > length-1 {
			return i + 1
		}
	}

	return num
}
