package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type head struct {
	x int
	y int
}

type tail struct {
	x int
	y int
}

type segment struct {
	x int
	y int
}

func main() {
	// log.SetLevel(log.DebugLevel)

	filePath := "day-9/input.txt"

	log.Print("First Star: ", firstStar(filePath))
	log.Print("Second Star: ", secondStar(filePath))
}

func firstStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	h := head{x: 0, y: 0}
	t := tail{x: 0, y: 0}
	tailVisited := make(map[tail]bool)
	tailVisited[t] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		instructions := strings.Split(line, " ")
		switch string(instructions[0]) {
		case "R":
			// log.Debug("Move right")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				h.x++
				switch {
				// Tail touching head
				case (h.y-1 <= t.y && t.y <= h.y+1) && (h.x-1 <= t.x && t.x <= h.x+1):
					log.Debug(t, h, "continue")
					continue
				// Tail on same column
				case t.x == h.x:
					log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
				// Tail on same row
				case t.y == h.y:
					if t.x == h.x-2 {
						t.x++
						tailVisited[t] = true
					} else {
						log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
					}
				default:
					if t.y < h.y {
						t.x++
						t.y++
						tailVisited[t] = true
						continue
					} else if t.y > h.y {
						t.x++
						t.y--
						tailVisited[t] = true
						continue
					}
					tailVisited[t] = true
					log.Fatalf("Solve diagonal situation? Tail [%v] Head [%v]", t, h)
				}
			}
		case "L":
			// log.Debug("Move left")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				h.x--
				switch {
				// Tail touching head
				case (h.y-1 <= t.y && t.y <= h.y+1) && (h.x-1 <= t.x && t.x <= h.x+1):
					log.Debug(t, h, "continue")
					continue
				// Tail on same column
				case t.x == h.x:
					log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
				// Tail on same row
				case t.y == h.y:
					if t.x == h.x+2 {
						t.x--
						tailVisited[t] = true
					} else {
						log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
					}
				default:
					if t.y < h.y {
						t.x--
						t.y++
						tailVisited[t] = true
						continue
					} else if t.y > h.y {
						t.x--
						t.y--
						tailVisited[t] = true
						continue
					}
					tailVisited[t] = true
					log.Fatalf("Solve diagonal situation? Tail [%v] Head [%v]", t, h)
				}
			}
		case "U":
			// log.Debug("Move up")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				h.y++
				switch {
				// Tail touching head
				case (h.y-1 <= t.y && t.y <= h.y+1) && (h.x-1 <= t.x && t.x <= h.x+1):
					log.Debug(t, h, "continue")
					continue
				// Tail on same row
				case t.y == h.y:
					log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
				// Tail on same column
				case t.x == h.x:
					if t.y == h.y-2 {
						t.y++
						tailVisited[t] = true
					} else {
						log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
					}
				default:
					if t.x < h.x {
						t.x++
						t.y++
						tailVisited[t] = true
						continue
					} else if t.x > h.x {
						t.x--
						t.y++
						tailVisited[t] = true
						continue
					}
					tailVisited[t] = true
					log.Fatalf("Solve diagonal situation? Tail [%v] Head [%v]", t, h)
				}
			}
		case "D":
			// log.Debug("Move down")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				h.y--
				switch {
				// Tail touching head
				case (h.y-1 <= t.y && t.y <= h.y+1) && (h.x-1 <= t.x && t.x <= h.x+1):
					log.Debug(t, h, "continue")
					continue
				// Tail on same row
				case t.y == h.y:
					log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
				// Tail on same column
				case t.x == h.x:
					if t.y == h.y+2 {
						t.y--
						tailVisited[t] = true
					} else {
						log.Fatalf("How can this happen without it beeing in proximity? Tail [%v] Head [%v]", t, h)
					}
				default:
					if t.x < h.x {
						t.x++
						t.y--
						tailVisited[t] = true
						continue
					} else if t.x > h.x {
						t.x--
						t.y--
						tailVisited[t] = true
						continue
					}
					tailVisited[t] = true
					log.Fatalf("Solve diagonal situation? Tail [%v] Head [%v]", t, h)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return len(tailVisited)
}

var segments []segment

func secondStar(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Debug("Len of segments array: ", len(segments))
	for i := 0; i < 10; i++ {
		segments = append(segments, segment{})
		segments[i].x = 0
		segments[i].y = 0
	}
	log.Debug("Len of segments array: ", len(segments))

	lastSegmentVisited := make(map[segment]bool)
	lastSegmentVisited[segments[len(segments)-1]] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		instructions := strings.Split(line, " ")
		switch string(instructions[0]) {
		case "R":
			// log.Debug("Move right")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				segments[0].x++
				log.Debug(i + 1)
				for index := range segments {
					if index != 0 {
						log.Debug("Segment No:", (index + 1))

						if !isWithinRange(segments[index-1], segments[index]) {
							segments[index] = moveSegment(segments[index-1], segments[index])
							log.Debug(segments[index])
						}
					}
				}
				lastSegmentVisited[segments[len(segments)-1]] = true
			}
		case "L":
			// log.Debug("Move left")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				segments[0].x--
				log.Debug(i + 1)
				for index := range segments {
					if index != 0 {
						log.Debug("Segment No:", (index + 1))

						if !isWithinRange(segments[index-1], segments[index]) {
							segments[index] = moveSegment(segments[index-1], segments[index])
							log.Debug(segments[index])
						}
					}
				}
				lastSegmentVisited[segments[len(segments)-1]] = true
			}
		case "U":
			// log.Debug("Move up")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				segments[0].y++
				log.Debug(i + 1)
				for index := range segments {
					if index != 0 {
						log.Debug("Segment No:", (index + 1))

						if !isWithinRange(segments[index-1], segments[index]) {
							segments[index] = moveSegment(segments[index-1], segments[index])
							log.Debug(segments[index])
						}
					}
				}
				lastSegmentVisited[segments[len(segments)-1]] = true
			}
		case "D":
			// log.Debug("Move down")
			num, _ := strconv.Atoi(string(instructions[1]))
			for i := 0; i < num; i++ {
				segments[0].y--
				log.Debug(i + 1)
				for index := range segments {
					if index != 0 {
						log.Debug("Segment No:", (index + 1))
						if !isWithinRange(segments[index-1], segments[index]) {
							segments[index] = moveSegment(segments[index-1], segments[index])
							log.Debug(segments[index])
						}
					}
				}
				lastSegmentVisited[segments[len(segments)-1]] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return len(lastSegmentVisited)
}

func isWithinRange(leader, follower segment) bool {
	positionsWithinRange := []segment{
		follower,
		{x: follower.x, y: follower.y + 1},
		{x: follower.x + 1, y: follower.y + 1},
		{x: follower.x + 1, y: follower.y},
		{x: follower.x + 1, y: follower.y - 1},
		{x: follower.x, y: follower.y - 1},
		{x: follower.x - 1, y: follower.y - 1},
		{x: follower.x - 1, y: follower.y},
		{x: follower.x - 1, y: follower.y + 1},
	}

	for i := 0; i < len(positionsWithinRange); i++ {
		if leader.x == positionsWithinRange[i].x && leader.y == positionsWithinRange[i].y {
			return true
		}
	}

	return false
}

func moveSegment(leader, follower segment) segment {
	log.Debug(leader, follower)

	positionsWithinRange := []segment{
		{x: follower.x + 2, y: follower.y + 1},
		{x: follower.x + 2, y: follower.y},
		{x: follower.x + 2, y: follower.y - 1},

		{x: follower.x - 2, y: follower.y + 1},
		{x: follower.x - 2, y: follower.y},
		{x: follower.x - 2, y: follower.y - 1},

		{x: follower.x + 1, y: follower.y + 2},
		{x: follower.x, y: follower.y + 2},
		{x: follower.x - 1, y: follower.y + 2},

		{x: follower.x + 1, y: follower.y - 2},
		{x: follower.x, y: follower.y - 2},
		{x: follower.x - 1, y: follower.y - 2},

		{x: follower.x + 2, y: follower.y + 2},
		{x: follower.x + 2, y: follower.y - 2},
		{x: follower.x - 2, y: follower.y + 2},
		{x: follower.x - 2, y: follower.y - 2},
	}

	switch {
	case leader.x == positionsWithinRange[0].x && leader.y == positionsWithinRange[0].y:
		return segment{x: follower.x + 1, y: follower.y + 1}
	case leader.x == positionsWithinRange[1].x && leader.y == positionsWithinRange[1].y:
		return segment{x: follower.x + 1, y: follower.y}
	case leader.x == positionsWithinRange[2].x && leader.y == positionsWithinRange[2].y:
		return segment{x: follower.x + 1, y: follower.y - 1}

	case leader.x == positionsWithinRange[3].x && leader.y == positionsWithinRange[3].y:
		return segment{x: follower.x - 1, y: follower.y + 1}
	case leader.x == positionsWithinRange[4].x && leader.y == positionsWithinRange[4].y:
		return segment{x: follower.x - 1, y: follower.y}
	case leader.x == positionsWithinRange[5].x && leader.y == positionsWithinRange[5].y:
		return segment{x: follower.x - 1, y: follower.y - 1}

	case leader.x == positionsWithinRange[6].x && leader.y == positionsWithinRange[6].y:
		return segment{x: follower.x + 1, y: follower.y + 1}
	case leader.x == positionsWithinRange[7].x && leader.y == positionsWithinRange[7].y:
		return segment{x: follower.x, y: follower.y + 1}
	case leader.x == positionsWithinRange[8].x && leader.y == positionsWithinRange[8].y:
		return segment{x: follower.x - 1, y: follower.y + 1}

	case leader.x == positionsWithinRange[9].x && leader.y == positionsWithinRange[9].y:
		return segment{x: follower.x + 1, y: follower.y - 1}
	case leader.x == positionsWithinRange[10].x && leader.y == positionsWithinRange[10].y:
		return segment{x: follower.x, y: follower.y - 1}
	case leader.x == positionsWithinRange[11].x && leader.y == positionsWithinRange[11].y:
		return segment{x: follower.x - 1, y: follower.y - 1}

	case leader.x == positionsWithinRange[12].x && leader.y == positionsWithinRange[12].y:
		return segment{x: follower.x + 1, y: follower.y + 1}
	case leader.x == positionsWithinRange[13].x && leader.y == positionsWithinRange[13].y:
		return segment{x: follower.x + 1, y: follower.y - 1}
	case leader.x == positionsWithinRange[14].x && leader.y == positionsWithinRange[14].y:
		return segment{x: follower.x - 1, y: follower.y + 1}
	case leader.x == positionsWithinRange[15].x && leader.y == positionsWithinRange[15].y:
		return segment{x: follower.x - 1, y: follower.y - 1}
	}

	log.Fatal("Cant Find leader", leader, follower)

	return segment{}
}
