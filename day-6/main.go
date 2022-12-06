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

func distinctLetters(s string) bool {
	letters := make(map[rune]bool)
	for _, letter := range s {
		letters[letter] = true
	}
	return len(letters) == len(s)
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
		if distinctLetters(checkLine) {
			return i + 1
		}
	}

	return num
}
