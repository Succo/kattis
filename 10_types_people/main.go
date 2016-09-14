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
	zoneMap := make(map[int]int)

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
			convex, zoneMap = propagate(area, convex, zoneMap, r1, c1, r2, c2, zone)
			_, maped := zoneMap[zone]
			if !maped {
				zoneMap[zone] = zone
			}
			zone++
		}
		if convex[r2][c2] == 0 {
			convex, zoneMap = propagate(area, convex, zoneMap, r2, c2, r1, c1, zone)
			_, maped := zoneMap[zone]
			if !maped {
				zoneMap[zone] = zone
			}
			zone++
		}
		// for _, line := range convex {
		// 	fmt.Println(line)
		// }
		// fmt.Println(zoneMap)
		// fmt.Println(convex[r1][c1])
		// fmt.Println(convex[r2][c2])
		if zoneMap[convex[r1][c1]] != zoneMap[convex[r2][c2]] {
			fmt.Println("neither")
		} else if area[r1][c1] == 0 {
			fmt.Println("binary")
		} else {
			fmt.Println("decimal")
		}
	}
}

func updateMap(zoneMap map[int]int, zone1 int, zone2 int) map[int]int {
	_, maped1 := zoneMap[zone1]
	_, maped2 := zoneMap[zone2]
	if !maped1 && maped2 {
		zoneMap[zone1] = zoneMap[zone2]
	} else if maped1 && !maped2 {
		zoneMap[zone2] = zoneMap[zone1]
	} else if maped1 && maped2 {
		for zone, realZone := range zoneMap {
			if realZone == zoneMap[zone2] {
				zoneMap[zone] = zoneMap[zone1]
			}
		}
	}
	return zoneMap
}

func propagate(area [][]int, convex [][]int, zoneMap map[int]int, startI int, startJ int, goalI int, goalJ int, zone int) ([][]int, map[int]int) {
	if convex[startI][startJ] != 0 {
		return convex, zoneMap
	}
	convex[startI][startJ] = zone
	if startI+1 < len(area) && area[startI][startJ] == area[startI+1][startJ] {
		if convex[startI+1][startJ] != 0 {
			zoneMap = updateMap(zoneMap, zone, convex[startI+1][startJ])
		}
		convex, zoneMap = propagate(area, convex, zoneMap, startI+1, startJ, goalI, goalJ, zone)
		if convex[goalI][goalJ] != 0 {
			return convex, zoneMap
		}
	}
	if startI-1 >= 0 && area[startI][startJ] == area[startI-1][startJ] {
		if convex[startI-1][startJ] != 0 {
			zoneMap = updateMap(zoneMap, zone, convex[startI-1][startJ])
		}
		convex, zoneMap = propagate(area, convex, zoneMap, startI-1, startJ, goalI, goalJ, zone)
		if convex[goalI][goalJ] != 0 {
			return convex, zoneMap
		}
	}
	if startJ+1 < len(area[startI]) && area[startI][startJ] == area[startI][startJ+1] {
		if convex[startI][startJ+1] != 0 {
			zoneMap = updateMap(zoneMap, zone, convex[startI][startJ+1])
		}
		convex, zoneMap = propagate(area, convex, zoneMap, startI, startJ+1, goalI, goalJ, zone)
		if convex[goalI][goalJ] != 0 {
			return convex, zoneMap
		}
	}
	if startJ-1 >= 0 && area[startI][startJ] == area[startI][startJ-1] {
		if convex[startI][startJ-1] != 0 {
			zoneMap = updateMap(zoneMap, zone, convex[startI][startJ-1])
		}
		convex, zoneMap = propagate(area, convex, zoneMap, startI, startJ-1, goalI, goalJ, zone)
		if convex[goalI][goalJ] != 0 {
			return convex, zoneMap
		}
	}
	return convex, zoneMap
}
