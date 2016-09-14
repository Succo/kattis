package main

import (
	"fmt"
)

const (
	mod = 1000000007
)

func getPermutSum(input string, numberOfOne int, numberOfPermut int) int {
	if len(input) == 0 {
		return 0
	}
	if input[0] == '0' {
		case0 := numberOfOne + getPermutSum(input[1:], numberOfOne, numberOfPermut+numberOfOne)
		return case0
	}
	if input[0] == '1' {
		case1 := getPermutSum(input[1:], numberOfOne+1, numberOfPermut)
		return case1
	}
	if input[0] == '?' {
		if0 := numberOfOne + getPermutSum(input[1:], numberOfOne, numberOfPermut+numberOfOne)
		if1 := getPermutSum(input[1:], numberOfOne+1, numberOfPermut)
		return numberOfPermut + if0 + if1
	}
	panic("invalid character")
}

func main() {
	var input string
	fmt.Scanln(&input)
	sum := getPermutSum(input, 0, 0)
	fmt.Printf("%d\n", sum%mod)
}
