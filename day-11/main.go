package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"sort"
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
	log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)

	filePath := "day-11/input_test.txt"

	getStar(filePath)
	arr := []int{}
	for i, m := range monkeys {
		log.Infof("Monkey %d inspected %d items", i, m.itemsInspected)
		arr = append(arr, m.itemsInspected)
	}

	sort.Ints(arr)

	// Reverse
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	monkeyBusiness := arr[0] * arr[1]

	fmt.Println("Monkey Business:", monkeyBusiness)
}

func getStar(filePath string) {
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

	commonMod := monkeys[0].divideBy
	for i := 1; i < len(monkeys); i++ {
		log.Debug(commonMod)
		commonMod *= monkeys[i].divideBy
	}
	log.Debug(commonMod)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rounds := 10000

	for i := 0; i < rounds; i++ {
		log.Debugln("Round:", i+1)
		for j := 0; j < len(monkeys); j++ {
			for _, item := range monkeys[j].items {
				log.Debug("Item: ", item)
				//Inspect
				n := 0
				ops := strings.Split(monkeys[j].operation, " ")

				switch ops[len(ops)-1] {
				case "old":
					n = item
				default:
					n, err = strconv.Atoi(string(ops[len(ops)-1]))
					if err != nil {
						log.Fatal(err)
					}
				}

				o := ops[len(ops)-2]
				switch o {
				case "*":
					monkeys[j].items[0] = item * n
				case "+":
					monkeys[j].items[0] = item + n
				}

				// Reduce stress level

				// Star 1
				//monkeys[j].items[0] = int(math.Floor(float64(monkeys[j].items[0]) / 3))

				// Star 2
				monkeys[j].items[0] = monkeys[j].items[0] % commonMod

				//Check division
				// isDivisible := monkeys[j].items[0]%monkeys[j].divideBy == 0
				isDivisible := monkeys[j].items[0] % monkeys[j].divideBy
				log.Debugln(monkeys[j].items[0], " % ", big.NewInt(int64(monkeys[j].divideBy)), " = ", isDivisible)

				//Throw to next monkey
				nextMonkey := 0
				if isDivisible == 0 {
					nextMonkey, _ = strconv.Atoi(string(monkeys[j].ifTrue[len(monkeys[j].ifTrue)-1]))
				} else {
					nextMonkey, _ = strconv.Atoi(string(monkeys[j].ifFalse[len(monkeys[j].ifFalse)-1]))
				}
				log.Debugln("Thrown to monkey number ", nextMonkey)
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
		num, _ := strconv.Atoi(string(item))
		nums = append(nums, num)
	}

	arr := strings.Split(monkeyArr[3], " ")
	divisibleBy, _ := strconv.Atoi(arr[len(arr)-1])

	monkeys = append(monkeys, Monkey{items: nums, operation: monkeyArr[2], divideBy: divisibleBy, ifTrue: monkeyArr[4], ifFalse: monkeyArr[5]})

}
