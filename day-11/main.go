package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Monkey struct {
	items          []int
	operation      string
	divideBy       int
	ifTrue         string
	ifFalse        string
	itemsInspected int
}

var monkeys []Monkey = []Monkey{}

func main() {
	//log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)

	filePath := "day-11/input_test.txt"

	firstStar(filePath)
	log.Print("First Star: ")
	for i, m := range monkeys {
		log.Infof("Monkey %d inspected %d items", i, m.itemsInspected)
	}

	secondStar(filePath)
	log.Printf("Second Star: \n")
}

func firstStar(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	monkeyArr := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug("Line: ", line)

		switch {
		case line != "":
			monkeyArr = append(monkeyArr, line)
		}
		if len(monkeyArr) == 6 {
			log.Debug("Parse Monkey")
			parseMonkey(monkeyArr)
			monkeyArr = []string{}
		}
	}

	for _, m := range monkeys {
		log.Debug(m)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rounds := 20

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, item := range monkeys[j].items {
				log.Debug(monkeys[j])
				//Inspect
				n := 0
				ops := strings.Split(monkeys[j].operation, " ")
				switch ops[len(ops)-1] {
				case "old":
					n = item
				default:
					n, _ = strconv.Atoi(ops[len(ops)-1])
				}

				o := ops[len(ops)-2]
				switch o {
				case "*":
					monkeys[j].items[0] = item * n
				case "+":
					monkeys[j].items[0] = item + n
				}

				// Divide
				f := math.Floor(float64(monkeys[j].items[0]) / 3)
				monkeys[j].items[0] = int(f)
				log.Debug(monkeys[j])

				//Check division
				isDivisible := monkeys[j].items[0]%monkeys[j].divideBy == 0

				//Throw to next monkey
				nextMonkey := 0
				if isDivisible {
					nextMonkey, _ = strconv.Atoi(string(monkeys[j].ifTrue[len(monkeys[j].ifTrue)-1]))
				} else {
					nextMonkey, _ = strconv.Atoi(string(monkeys[j].ifFalse[len(monkeys[j].ifFalse)-1]))
				}
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, monkeys[j].items[0])
				monkeys[j].items = monkeys[j].items[1:]

				//Add 1 to inspected items
				monkeys[j].itemsInspected++
			}
		}
	}
}

func parseMonkey(monkeyArr []string) {
	re := regexp.MustCompile("[0-9]+")
	items := re.FindAllString(monkeyArr[1], -1)
	nums := []int{}
	for _, item := range items {
		num, _ := strconv.Atoi(item)
		nums = append(nums, num)
	}

	arr := strings.Split(monkeyArr[3], " ")
	divisibleBy, _ := strconv.Atoi(arr[len(arr)-1])

	monkeys = append(monkeys, Monkey{items: nums, operation: monkeyArr[2], divideBy: divisibleBy, ifTrue: monkeyArr[4], ifFalse: monkeyArr[5]})

}

func secondStar(filePath string) {

}
