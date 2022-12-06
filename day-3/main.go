package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func main() {
	filePath := "day-3/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	items := ""

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		compartment1 := line[:len(line)/2]
		compartment1 = distinct(compartment1)
		// fmt.Println(compartment1)
		compartment2 := line[len(line)/2:]
		compartment2 = distinct(compartment2)
		// fmt.Println(compartment2)

		foundItem := false
		for _, item1 := range compartment1 {
			for _, item2 := range compartment2 {
				if item1 == item2 {
					items += string(item1)
					foundItem = true
					break
				}
			}
			if foundItem {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sumPriority(items)
}

func secondStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	items := ""
	packs := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		packs = append(packs, distinct(line))

		if len(packs) == 3 {
			items += getCommonItem(packs)
			packs = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sumPriority(items)
}

func distinct(s string) string {
	letters := make(map[rune]bool)
	res := ""
	for _, letter := range s {
		exists := letters[letter]
		if exists {
			continue
		}
		res += string(letter)
		letters[letter] = true
	}
	return res
}

func sumPriority(items string) int {
	sum := 0

	for _, item := range items {
		if unicode.IsUpper(item) {
			sum += int(item) - 38
		} else {
			sum += int(item) - 96
		}
	}

	return sum
}

func getCommonItem(packs []string) string {
	a := packs[0]
	b := packs[1]
	c := packs[2]
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	m := make(map[rune]uint8)
	for _, k := range a {
		m[k] |= (1 << 0)
	}
	for _, k := range b {
		m[k] |= (1 << 1)
	}
	for _, k := range c {
		m[k] |= (1 << 2)
	}

	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		c := v&(1<<2) != 0
		switch {
		case a && b && c:
			// fmt.Printf("Common item: %s\n", string(k))
			return string(k)
		}
	}

	return ""
}
