package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type TreeArray struct {
	rows []string
}

var filePath string = "day-8/input.txt"

func main() {
	treeStringArray := TreeArray{rows: []string{}}

	// Populate tree array
	treeStringArray.getTreeStringArray()
	// for _, row := range treeStringArray.rows {
	// 	fmt.Println(row)
	// }

	firstStarResult := treeStringArray.visibleTreesCount()

	log.Print("First Star: ", firstStarResult)
	log.Print("Second Star: ", treeStringArray.bestTree())
}

func (t *TreeArray) getTreeStringArray() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t.rows = append(t.rows, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (t *TreeArray) getVerticalRow(index int) string {
	row := ""
	for i := 0; i < len(t.rows); i++ {
		row += string(t.rows[i][index])
	}
	return row
}

func (t *TreeArray) visibleTreesCount() int {
	count := 0

	// Add outer edge
	count += (len(t.rows) * 2) + (len(t.rows[0]) * 2) - 4
	fmt.Println("Trees visible on the outer edge: ", count)

	for y := 1; y < len(t.rows)-1; y++ {
		for i := 1; i < len(t.rows[y])-1; i++ {
			// fmt.Println("=================")
			row := t.rows[y]
			// fmt.Println(row)

			height, _ := strconv.Atoi(string(t.rows[y][i]))
			visible1 := treeVisible(row, height, i)

			// fmt.Printf("horizontal - %v\n", visible1)

			row = t.getVerticalRow(i)
			height, _ = strconv.Atoi(string(t.rows[y][i]))
			// fmt.Println(row)
			visible2 := treeVisible(row, height, y)

			// fmt.Printf("vertical - %v\n", visible2)

			visible := (visible1 || visible2)

			// fmt.Printf("[%d] visibility: %v\n", height, visible)

			if visible {
				count++
			}
		}
	}

	return count
}

func (t *TreeArray) bestTree() int {
	count := 0

	for y := 1; y < len(t.rows)-1; y++ {
		for i := 1; i < len(t.rows[y])-1; i++ {
			sum := 0
			row := t.rows[y]

			height, _ := strconv.Atoi(string(t.rows[y][i]))

			//LOOK LEFT
			left := 0
			for j := i - 1; j >= 0; j-- {
				tree, _ := strconv.Atoi(string(row[j]))
				if tree >= height {
					left++
					break
				}
				left++
			}

			//LOOK RIGHT
			right := 0
			for j := i + 1; j <= len(row)-1; j++ {
				tree, _ := strconv.Atoi(string(row[j]))
				if tree >= height {
					right++
					break
				}
				right++
			}

			row = t.getVerticalRow(i)

			//LOOK UP
			up := 0
			for j := y - 1; j >= 0; j-- {
				tree, _ := strconv.Atoi(string(row[j]))
				if tree >= height {
					up++
					break
				}
				up++
			}

			//LOOK DOWN
			down := 0
			for j := y + 1; j <= len(row)-1; j++ {
				tree, _ := strconv.Atoi(string(row[j]))
				if tree >= height {
					down++
					break
				}
				down++
			}

			sum = up * down * left * right
			fmt.Printf("[%d]\n", height)
			fmt.Printf("up %d, left %d, right %d, down %d\n", up, left, right, down)
			fmt.Println(sum)

			if count < sum {
				count = sum
			}
		}
	}

	return count
}

func treeVisible(row string, height int, index int) bool {
	visible1 := true
	visible2 := true
	for i, tree := range row {
		if i == index {
			// fmt.Printf("[%d]", height)
			continue
		}
		treeHeight, _ := strconv.Atoi(string(tree))
		// fmt.Print(treeHeight)
		if treeHeight >= height {
			if i < index {
				visible1 = false
			}
			if i > index {
				visible2 = false
			}
		}
	}
	// fmt.Println()
	// fmt.Printf("\nVISIBILITY")
	// fmt.Printf("\n%v<-o->%v\n", visible1, visible2)
	return visible1 || visible2
}
