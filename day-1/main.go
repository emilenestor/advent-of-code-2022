package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves := []int{}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, sum)
			sum = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Print(err)
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	log.Print(elves[0] + elves[1] + elves[2])
}
