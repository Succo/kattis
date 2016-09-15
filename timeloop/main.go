package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)
	i, _ := strconv.Atoi(input)
	for j := 1; j <= i; j++ {
		fmt.Printf("%d Abracadabra\n", j)
	}
}
