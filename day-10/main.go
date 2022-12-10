package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)

	filePath := "day-10/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	secondStar(filePath)
	log.Printf("Second Star: \n")
	for i := 0; i < len(screen); i++ {
		fmt.Println(screen[i])
	}
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 1
	cycle := 1
	signalStrength := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug("Line: ", line)
		splitLine := strings.Split(line, " ")
		switch {
		case len(splitLine) == 1:
			log.Debugf("Read: [%s]", splitLine[0])
			log.Debugf("Add tick: [%d -> %d]", cycle, cycle+1)
			cycle++
			if calculateSignalStrenght(cycle) {
				log.Debugf("CYCLE [%d] - Calculating signal strength", cycle)
				signalStrength += cycle * x
				log.Debugf("SIGNAL STRENGTH [%d]", signalStrength)
			}
		case len(splitLine) == 2:
			log.Debugf("Read: [%s %s]", splitLine[0], splitLine[1])

			for i := 0; i < 2; i++ {
				log.Debugf("Add tick: [%d -> %d]", cycle, cycle+1)
				cycle++
				if i == 1 {
					log.Debugf("Before adding [%s] to X, X equals [%d]", splitLine[1], x)
					addx, _ := strconv.Atoi(string(splitLine[1]))
					x += addx
					log.Debugf("After adding [%d] to X, X equals [%d]", addx, x)
				}
				if calculateSignalStrenght(cycle) {
					log.Debugf("CYCLE [%d] - Calculating signal strength", cycle)
					signalStrength += cycle * x
					log.Debugf("SIGNAL STRENGTH [%d]", signalStrength)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return signalStrength
}

func calculateSignalStrenght(cycle int) bool {
	firstCalculationCycle := 20
	if cycle == firstCalculationCycle {
		return true
	}
	if (cycle-firstCalculationCycle)%40 == 0 {
		return true
	}
	return false
}

var screen []string

func secondStar(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	screen = make([]string, 6)

	x := 1
	cycle := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug("Line: ", line)
		splitLine := strings.Split(line, " ")
		switch {
		case len(splitLine) == 1:
			writeToScreen(cycle, x)
			log.Debugf("Read: [%s]", splitLine[0])
			log.Debugf("Add tick: [%d -> %d]", cycle, cycle+1)
			cycle++
		case len(splitLine) == 2:
			log.Debugf("Read: [%s %s]", splitLine[0], splitLine[1])

			for i := 0; i < 2; i++ {
				writeToScreen(cycle, x)
				log.Debugf("Add tick: [%d -> %d]", cycle, cycle+1)
				cycle++

				if i == 1 {
					log.Infof("Before adding [%s] to X, X equals [%d]", splitLine[1], x)
					addx, _ := strconv.Atoi(string(splitLine[1]))
					x += addx
					log.Infof("After adding [%d] to X, X equals [%d]", addx, x)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func writeToScreen(cycle, x int) {
	log.Debugf("Writing to pixel [%d]", cycle)
	log.Debugf("X are in positions [%d, %d and %d]", x, x+1, x+2)

	log.Infof("CYCLE [%d]", cycle)
	if cycle > 40 {
		if cycle%40 != 0 {
			cycle %= 40
		} else {
			cycle = 40
		}
	}
	overlap := cycle == x || cycle == x+1 || cycle == x+2
	log.Infof("Writing to pixel [%d]", cycle)
	log.Infof("X are in positions [%d, %d and %d]", x, x+1, x+2)
	log.Infof("They overlap [%v]", overlap)

	for i := 0; i < len(screen); i++ {
		if len(screen[i]) == 40 {
			continue
		}
		switch {
		case cycle >= x && cycle <= x+2:
			screen[i] += "#"
			return
		default:
			screen[i] += "."
			return
		}
	}
}
