package main

import (
	// "bytes"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"time"
)

// contiguous is a contiguous set of identical cases
// denoted by it's starting value and end value
type contiguous struct {
	x, y  int
	value int
}

func main() {
	var start, step1 time.Time
	timing := false
	if timing {
		start = time.Now()
	}
	var r, c int
	fmt.Scanln(&r, &c)
	area := make([][]contiguous, r)
	for i := 0; i < r; i++ {
		area[i] = []contiguous{}
		var line string
		fmt.Scanln(&line)
		presentZone := contiguous{value: 2}
		for index, j := range line {
			value, _ := strconv.Atoi(string(j))
			if presentZone.value == value {
				continue
			} else {
				if presentZone.value == 2 {
					presentZone = contiguous{x: index, value: value}
				} else {
					presentZone.y = index - 1
					area[i] = append(area[i], presentZone)
					presentZone = contiguous{x: index, value: value}
				}
			}
		}
		presentZone.y = c - 1
		area[i] = append(area[i], presentZone)
	}
	var n int
	fmt.Scanln(&n)
	if timing {
		fmt.Printf("Parsing %s \n", time.Since(start))
		step1 = time.Now()
	}

	convex := make([][]contiguous, len(area))
	for i, line := range area {
		convex[i] = make([]contiguous, len(line))
		for j, zone := range line {
			convex[i][j].x = zone.x
			convex[i][j].y = zone.y
		}
	}
	zone := 1

	for i := 0; i < n; i++ {
		var r1, c1, r2, c2 int
		fmt.Scanln(&r1, &c1, &r2, &c2)
		r1 = r1 - 1
		c1 = c1 - 1
		r2 = r2 - 1
		c2 = c2 - 1
		if getValue(area, r1, c1) != getValue(area, r2, c2) {
			if !timing {
				fmt.Println("neither")
			}
			continue
		}
		if getValue(convex, r1, c1) == 0 && getValue(convex, r2, c2) == 0 {
			index := getIndex(area, r1, c1)
			convex = propagate(area, convex, r1, index, zone)
			zone++
		}
		if getValue(convex, r1, c1) != getValue(convex, r2, c2) {
			if !timing {
				fmt.Println("neither")
			}
		} else if getValue(area, r1, c1) == 0 {
			if !timing {
				fmt.Println("binary")
			}
		} else {
			if !timing {
				fmt.Println("decimal")
			}
		}
	}
	if timing {
		fmt.Printf("Pathing %s \n", time.Since(step1))
	}
	// printArea(area, convex)
}

func printArea(area [][]contiguous, convex [][]contiguous) {
	colors := []color.Attribute{color.FgWhite, color.FgRed, color.FgGreen, color.FgYellow, color.FgBlue, color.FgMagenta, color.FgCyan}
	for i, line := range area {
		if i < 10 {
			fmt.Print("0")
		}
		fmt.Printf("%d ", i)
		for j, subZone := range line {
			for index := subZone.x; index <= subZone.y; index++ {
				zoneColor := color.New(colors[convex[i][j].value%7])
				zoneColor.Print(strconv.Itoa(subZone.value))

			}
		}
		fmt.Print("   ")
		for j, subZone := range convex[i] {
			for index := subZone.x; index <= subZone.y; index++ {
				zoneColor := color.New(colors[(area[i][j].value+1)%7])
				// big hack
				if subZone.value < 10 {
					zoneColor.Print("0")
				}
				zoneColor.Print(strconv.Itoa(subZone.value))

			}
		}
		fmt.Println("")
	}
}

func getValue(area [][]contiguous, i int, j int) int {
	for _, zone := range area[i] {
		if zone.x <= j && j <= zone.y {
			return zone.value
		}
	}
	return 0
}

func getIndex(area [][]contiguous, i int, j int) int {
	for index, zone := range area[i] {
		if zone.x <= j && j <= zone.y {
			return index
		}
	}
	panic(fmt.Sprintf("getValue wrong dimensions %d, %d", i, j))
}

// getZones return all the zones in the line above and below
func getZones(area [][]contiguous, lineIndex int, zone contiguous) ([]int, []int) {
	bellow := []int{}
	above := []int{}
	if lineIndex-1 >= 0 {
		for index, subLine := range area[lineIndex-1] {
			if zone.value == subLine.value &&
				((zone.x <= subLine.x && subLine.x <= zone.y) || (zone.x <= subLine.y && subLine.y <= zone.y) ||
					(subLine.x <= zone.x && zone.x <= subLine.y) || (subLine.x <= zone.y && zone.y <= subLine.y)) {
				bellow = append(bellow, index)
			}
		}
	}
	if lineIndex+1 < len(area) {
		for index, subLine := range area[lineIndex+1] {
			if zone.value == subLine.value &&
				((zone.x <= subLine.x && subLine.x <= zone.y) || (zone.x <= subLine.y && subLine.y <= zone.y) ||
					(subLine.x <= zone.x && zone.x <= subLine.y) || (subLine.x <= zone.y && zone.y <= subLine.y)) {
				above = append(above, index)
			}
		}
	}
	return bellow, above
}

func propagate(area [][]contiguous, convex [][]contiguous, lineIndex int, zoneIndex int, color int) [][]contiguous {
	// If the zone is mapped, ignore it
	if convex[lineIndex][zoneIndex].value != 0 {
		return convex
	}
	// Otherwise color the zone
	convex[lineIndex][zoneIndex].value = color
	bellow, above := getZones(area, lineIndex, area[lineIndex][zoneIndex])
	for _, index := range above {
		convex = propagate(area, convex, lineIndex+1, index, color)
	}
	for _, index := range bellow {
		convex = propagate(area, convex, lineIndex-1, index, color)
	}
	return convex
}
