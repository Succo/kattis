package main

import (
	"fmt"
	"github.com/fatih/color"
)

func printArea(area []string, convex [][]int) {
	colors := []color.Attribute{color.FgRed, color.FgGreen, color.FgYellow, color.FgBlue, color.FgMagenta, color.FgCyan}
	for i, line := range area {
		if i < 10 {
			fmt.Print("0")
		}
		fmt.Printf("%d ", i)
		for j, value := range line {
			var zoneColor *color.Color
			if convex[i][j] == 0 {
				zoneColor = color.New(color.FgWhite)
			} else {
				zoneColor = color.New(colors[convex[i][j]%6])
			}
			zoneColor.Print(string(value))

		}
		// fmt.Print("   ")
		// for j, subZone := range convex[i] {
		// 	for index := subZone.x; index <= subZone.y; index++ {
		// 		zoneColor := color.New(colors[(area[i][j].value+1)%7])
		// 		// big hack
		// 		if subZone.value < 10 {
		// 			zoneColor.Print("0")
		// 		}
		// 		zoneColor.Print(strconv.Itoa(subZone.value))
		//
		// 	}
		// }
		fmt.Println("")
	}
}
