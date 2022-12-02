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

	lose = "X"
	draw = "Y"
	win  = "Z"
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

	log.Print("Second Star: ", sum)
}
