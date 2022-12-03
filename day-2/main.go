package main

import (
	"bufio"
	"log"
	"os"
)

const (
	aRock    = "A"
	bPaper   = "B"
	cScissor = "C"

	xRock    = "X"
	yPaper   = "Y"
	zScissor = "Z"

	lose = "X"
	draw = "Y"
	win  = "Z"
)

func main() {
	filePath := "day-2/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch string(line[2]) {
		case xRock:
			sum += 1
			switch string(line[0]) {
			case aRock:
				sum += 3
			case bPaper:
				sum += 0
			case cScissor:
				sum += 6
			}
		case yPaper:
			sum += 2
			switch string(line[0]) {
			case aRock:
				sum += 6
			case bPaper:
				sum += 3
			case cScissor:
				sum += 0
			}
		case zScissor:
			sum += 3
			switch string(line[0]) {
			case aRock:
				sum += 0
			case bPaper:
				sum += 6
			case cScissor:
				sum += 3
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func secondStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch string(line[2]) {
		case lose:
			switch string(line[0]) {
			case aRock:
				sum += 0 + 3
			case bPaper:
				sum += 0 + 1
			case cScissor:
				sum += 0 + 2
			}
		case draw:
			switch string(line[0]) {
			case aRock:
				sum += 3 + 1
			case bPaper:
				sum += 3 + 2
			case cScissor:
				sum += 3 + 3
			}
		case win:
			switch string(line[0]) {
			case aRock:
				sum += 6 + 2
			case bPaper:
				sum += 6 + 3
			case cScissor:
				sum += 6 + 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
