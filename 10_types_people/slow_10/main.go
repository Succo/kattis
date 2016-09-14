package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const ()

func main() {
	var start, step1 time.Time
	timing := false
	if timing {
		start = time.Now()
	}
	in := bufio.NewReader(os.Stdin)
	var r, c int
	fmt.Fscanln(in, &r, &c)
	area := make([]string, r)
	for i := 0; i < r; i++ {
		line, err := in.ReadString('\n')
		if err != nil {
			panic("error reading stdin")
		}
		area[i] = line[:len(line)-1]
	}
	var n int
	fmt.Fscanln(in, &n)
	if timing {
		fmt.Printf("Parsing %s \n", time.Since(start))
		step1 = time.Now()
	}

	convex := make([][]int, len(area))

	for i := 0; i < r; i++ {
		convex[i] = make([]int, c)
	}
	zone := 1

	for i := 0; i < n; i++ {
		var r1, c1, r2, c2 int
		fmt.Fscanln(in, &r1, &c1, &r2, &c2)
		r1 = r1 - 1
		c1 = c1 - 1
		r2 = r2 - 1
		c2 = c2 - 1
		fmt.Printf("%d ", i)
		if area[r1][c1] != area[r2][c2] {
			if !timing {
				fmt.Println("neither")
			}
			continue
		}
		if convex[r1][c1] == 0 && convex[r2][c2] == 0 {
			convex = propagate(area, convex, r1, c1, zone)
			zone++
		}
		if convex[r1][c1] != convex[r2][c2] {
			if !timing {
				fmt.Println("neither")
			}
		} else if area[r1][c1] == byte(48) {
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
	printArea(area, convex)
}

func propagate(area []string, convex [][]int, startI int, startJ int, zone int) [][]int {
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
