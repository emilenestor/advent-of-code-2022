package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	from int
	to   int
}

func main() {
	filePath := "day-4/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func convertToElves(line string) []Elf {
	firstSplit := strings.Split(line, ",")

	secondSplit := strings.Split(firstSplit[0], "-")

	from, _ := strconv.Atoi(string(secondSplit[0]))
	to, _ := strconv.Atoi(string(secondSplit[1]))

	elves := []Elf{}
	elves = append(elves, Elf{from: from, to: to})

	secondSplit = strings.Split(firstSplit[1], "-")

	from, _ = strconv.Atoi(string(secondSplit[0]))
	to, _ = strconv.Atoi(string(secondSplit[1]))

	elves = append(elves, Elf{from: from, to: to})

	return elves
}

func entireAreaOverlap(elves []Elf) bool {
	if len(elves) < 2 {
		return false
	}

	return (elves[0].from >= elves[1].from && elves[0].to <= elves[1].to) || (elves[1].from >= elves[0].from && elves[1].to <= elves[0].to)
}

func anyAreaOverlap(elves []Elf) bool {
	if len(elves) < 2 {
		return false
	}

	return ((elves[0].from >= elves[1].from && elves[0].from <= elves[1].to) ||
		(elves[0].to >= elves[1].from && elves[0].to <= elves[1].to) ||
		(elves[1].from >= elves[0].from && elves[1].from <= elves[0].to))
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
		// fmt.Println(line)

		elves := convertToElves(line)
		// fmt.Println(elves)

		// overlap := entireAreaOverlap(elves)
		// fmt.Println(overlap)

		if overlap := entireAreaOverlap(elves); overlap {
			num++
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

	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		elves := convertToElves(line)
		// fmt.Println(elves)

		// overlap := entireAreaOverlap(elves)
		// fmt.Println(overlap)

		if overlap := anyAreaOverlap(elves); overlap {
			num++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return num
}
