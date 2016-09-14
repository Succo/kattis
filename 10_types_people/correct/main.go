package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r, c int
	fmt.Scanln(&r, &c)
	area := make([][]int, r)
	for i := 0; i < r; i++ {
		area[i] = make([]int, c)
		var line string
		fmt.Scanln(&line)
		for j := 0; j < c; j++ {
			area[i][j], _ = strconv.Atoi(string(line[j]))
		}
	}
	var n int
	fmt.Scanln(&n)

	convex := make([][]int, len(area))
	for i := 0; i < len(area); i++ {
		convex[i] = make([]int, len(area[0]))
	}
	zone := 1

	for i := 0; i < n; i++ {
		var r1, c1, r2, c2 int
		fmt.Scanln(&r1, &c1, &r2, &c2)
		r1 = r1 - 1
		c1 = c1 - 1
		r2 = r2 - 1
		c2 = c2 - 1
		if area[r1][c1] != area[r2][c2] {
			fmt.Println("neither")
			continue
		}
		if convex[r1][c1] == 0 && convex[r2][c2] == 0 {
			convex = propagate(area, convex, r1, c1, zone)
			zone++
		}
		if convex[r1][c1] != convex[r2][c2] {
			fmt.Println("neither")
		} else if area[r1][c1] == 0 {
			fmt.Println("binary")
		} else {
			fmt.Println("decimal")
		}
	}
}

func propagate(area [][]int, convex [][]int, startI int, startJ int, zone int) [][]int {
	if convex[startI][startJ] != 0 {
		return convex
	}
	convex[startI][startJ] = zone
	if startI+1 < len(area) && area[startI][startJ] == area[startI+1][startJ] {
		convex = propagate(area, convex, startI+1, startJ, zone)
	}
	if startI-1 >= 0 && area[startI][startJ] == area[startI-1][startJ] {
		convex = propagate(area, convex, startI-1, startJ, zone)
	}
	if startJ+1 < len(area[startI]) && area[startI][startJ] == area[startI][startJ+1] {
		convex = propagate(area, convex, startI, startJ+1, zone)
	}
	if startJ-1 >= 0 && area[startI][startJ] == area[startI][startJ-1] {
		convex = propagate(area, convex, startI, startJ-1, zone)
	}
	return convex
}
