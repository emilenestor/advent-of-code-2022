package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "day-0/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
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

	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return num
}
