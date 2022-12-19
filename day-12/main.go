package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Coordinate struct {
	X int
	Y int
}

type square struct {
	Height    rune
	SteppedOn bool
}

var (
	heightMap map[Coordinate]square = make(map[Coordinate]square)
	start     Coordinate
	end       Coordinate
	steps     int = 0
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)

	filePath := "day-12/input_test.txt"

	getStar(filePath)

	fmt.Println("START")
	fmt.Println(start)
	fmt.Println("END")
	fmt.Println(end)
	printMap()
	fmt.Println("STEPS")
	fmt.Println(steps)
}

func getStar(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug("Line: ", line)

		for i, v := range line {
			height := v
			heightMap[Coordinate{X: i, Y: row}] = square{height, false}
			if height == 'S' {
				start = Coordinate{X: i, Y: row}
			} else if height == 'E' {
				end = Coordinate{X: i, Y: row}
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	calculateSteps()
}

func calculateSteps() {
	currentPosition := start
	for currentPosition != end {
		printMap()
		time.Sleep(time.Millisecond * 1)
		triedDirection := ""

		if !heightMap[Coordinate{X: currentPosition.X - 1, Y: currentPosition.Y}].SteppedOn && currentPosition.X > end.X {
			triedDirection = "left"
			if _, exists := heightMap[Coordinate{X: currentPosition.X - 1, Y: currentPosition.Y}]; !exists {
				continue
			}

			toPosition := Coordinate{currentPosition.X - 1, currentPosition.Y}
			if canMove("left", toPosition, currentPosition) {
				//move left
				heightMap[currentPosition] = square{'<', true}
				currentPosition.X--
				steps++
				continue
			}
		}
		if !heightMap[Coordinate{X: currentPosition.X + 1, Y: currentPosition.Y}].SteppedOn && currentPosition.X < end.X {
			triedDirection = "right"
			if _, exists := heightMap[Coordinate{X: currentPosition.X + 1, Y: currentPosition.Y}]; !exists {
				continue
			}

			toPosition := Coordinate{currentPosition.X + 1, currentPosition.Y}
			if canMove("right", toPosition, currentPosition) {
				//move right
				heightMap[currentPosition] = square{'>', true}
				currentPosition.X++
				steps++
				continue
			}
		}
		if !heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y - 1}].SteppedOn && currentPosition.Y > end.Y {
			triedDirection = "up"
			if _, exists := heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y - 1}]; !exists {
				continue
			}

			toPosition := Coordinate{currentPosition.X, currentPosition.Y - 1}
			if canMove("up", toPosition, currentPosition) {
				//move up
				heightMap[currentPosition] = square{'^', true}
				currentPosition.Y--
				steps++

				continue
			}
		}
		if !heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y + 1}].SteppedOn && currentPosition.Y < end.Y {
			triedDirection = "down"
			if _, exists := heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y + 1}]; !exists {
				continue
			}

			toPosition := Coordinate{currentPosition.X, currentPosition.Y + 1}
			if canMove("down", toPosition, currentPosition) {
				//move down
				heightMap[currentPosition] = square{'v', true}
				currentPosition.Y++
				steps++

				continue
			}
		}
		if triedDirection != "" {
			switch triedDirection {
			case "right":
				if !heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y + 1}].SteppedOn {
					//move down
					heightMap[currentPosition] = square{'v', true}
					currentPosition.Y++
					steps++

					continue

				} else if !heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y - 1}].SteppedOn {
					//move up
					heightMap[currentPosition] = square{'^', true}
					currentPosition.Y--
					steps++

					continue
				}
			case "left":
				if !heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y + 1}].SteppedOn {
					//move down
					heightMap[currentPosition] = square{'v', true}
					currentPosition.Y++
					steps++

					continue

				} else if !heightMap[Coordinate{X: currentPosition.X, Y: currentPosition.Y - 1}].SteppedOn {
					//move up
					heightMap[currentPosition] = square{'^', true}
					currentPosition.Y--
					steps++

					continue
				}
			case "up":
				if !heightMap[Coordinate{X: currentPosition.X - 1, Y: currentPosition.Y}].SteppedOn {
					//move left
					heightMap[currentPosition] = square{'<', true}
					currentPosition.X--
					steps++

					continue

				} else if !heightMap[Coordinate{X: currentPosition.X + 1, Y: currentPosition.Y}].SteppedOn {
					//move right
					heightMap[currentPosition] = square{'>', true}
					currentPosition.X++
					steps++

					continue
				}
			case "down":
				if !heightMap[Coordinate{X: currentPosition.X - 1, Y: currentPosition.Y}].SteppedOn {
					//move left
					heightMap[currentPosition] = square{'<', true}
					currentPosition.X--
					steps++

					continue

				} else if !heightMap[Coordinate{X: currentPosition.X + 1, Y: currentPosition.Y}].SteppedOn {
					//move right
					heightMap[currentPosition] = square{'>', true}
					currentPosition.X++
					steps++

					continue
				}
			}
		}
	}
}

func canMove(direction string, to, from Coordinate) bool {
	currentHeight := heightMap[from]
	if currentHeight.Height == 'S' {
		return true
	}

	s := heightMap[to]
	s2 := heightMap[from]
	if s.Height == 'E' && s2.Height != 'z' {
		return false
	} else if s.Height == 'E' && s2.Height == 'z' {
		return true
	}

	switch direction {
	case "up":
		if s.Height <= currentHeight.Height+1 && s.Height >= currentHeight.Height {
			return true
		}
	case "down":
		if s.Height <= currentHeight.Height+1 && s.Height >= currentHeight.Height {
			return true
		}
	case "left":
		if s.Height <= currentHeight.Height+1 && s.Height >= currentHeight.Height {
			return true
		}
	case "right":
		if s.Height <= currentHeight.Height+1 && s.Height >= currentHeight.Height {
			return true
		}
	}

	return false
}

func printMap() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(string(heightMap[Coordinate{j, i}].Height))
		}
		fmt.Print("\n")
	}
}
