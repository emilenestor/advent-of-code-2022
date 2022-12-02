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
)

func main() {
	file, err := os.Open("day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

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

	log.Print("First Star: ", sum)
}
