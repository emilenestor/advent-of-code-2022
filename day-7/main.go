package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := "day-7/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	// log.Print("Second Star: ", secondStar(filePath))
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileMap := make(map[string]int)
	dirMap := make(map[string]int)
	dirMap["/"] = 0

	path := "/"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		re := regexp.MustCompile("[0-9]+")

		cmdLine := strings.Split(line, " ")
		if string(cmdLine[0]) == "$" {
			switch string(cmdLine[1]) {
			case "cd":
				switch {
				case string(cmdLine[2]) == "..":
					split := strings.Split(path, "/")
					i := len(split[len(split)-1])
					if len(split) > 1 {
						i++
					}
					fmt.Println(path, " to ", path[:len(path)-i])
					path = path[:len(path)-i]
					if !strings.HasPrefix(path, "/") {
						path = fmt.Sprintf("%s%s", "/", path)
					}
				default:
					if !strings.HasPrefix(path, "/") {
						path = fmt.Sprintf("%s%s", "/", path)
					}
					if len(path) > 1 {
						path += "/" + string(cmdLine[2])
					} else if cmdLine[2] != "/" {
						path += string(cmdLine[2])
					}
				}
			}
		}

		if cmdLine[0] == "dir" {
			if path == "/" {
				dirMap[path+cmdLine[1]] = 0
			} else {
				dirMap[path+"/"+cmdLine[1]] = 0
			}
		}

		isNum := re.Match([]byte(string(cmdLine[0])))
		if isNum {
			size, _ := strconv.Atoi(string(cmdLine[0]))
			if path == "/" {
				fileMap[path+cmdLine[1]] += size
			} else {
				fileMap[path+"/"+cmdLine[1]] += size
			}
		}

		// IF file add to map[filepath]filesize
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for dir := range dirMap {
		for file, size := range fileMap {
			if strings.HasPrefix(file, dir) {
				dirMap[dir] += size
			}
		}
	}

	// Sum all values below 100000
	sum := 0
	for _, size := range dirMap {
		if size < 100000 {
			sum += size
		}
	}

	keys := make([]string, 0, len(dirMap))

	for key := range dirMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return dirMap[keys[i]] < dirMap[keys[j]]
	})
	required := (dirMap["/"] - (70000000 - 30000000))
	allFiles := dirMap["/"]
	fmt.Println("Space Required: ", required)
	fmt.Println("Main: ", allFiles)
	for _, k := range keys {
		fmt.Println("Directory size: ", dirMap[k])
		if required <= dirMap[k] {
			sum = dirMap[k]
			break
		}
	}

	return sum
}

func secondStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return num
}
