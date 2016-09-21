package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	mod = 1000000007
)

func getPermutSum(buff *bufio.Reader) (int, int, int) {
	b, _, _ := buff.ReadRune()
	if b == '\n' {
		return 0, 0, 1
	}
	if b == '0' {
		permuts, zeros, paths := getPermutSum(buff)
		return permuts, zeros + paths, paths
	}
	if b == '1' {
		permut, zeros, paths := getPermutSum(buff)
		return permut + zeros, zeros, paths
	}
	if b == '?' {
		permut, zeros, paths := getPermutSum(buff)
		return 2*permut + zeros, 2*zeros + paths, paths * 2
	}
	panic("invalid character")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	sum, _, _ := getPermutSum(in)
	fmt.Printf("%d\n", sum%mod)
}

func getOldPermutSum(input string, numberOfOne int, numberOfPermut int) int {
	if len(input) == 0 {
		return 0
	}
	// fmt.Println(input)
	if input[0] == '0' {
		case0 := numberOfOne + getOldPermutSum(input[1:], numberOfOne, numberOfPermut+numberOfOne)
		return case0
	}
	if input[0] == '1' {
		case1 := getOldPermutSum(input[1:], numberOfOne+1, numberOfPermut)
		return case1
	}
	if input[0] == '?' {
		if0 := numberOfOne + getOldPermutSum(input[1:], numberOfOne, numberOfPermut+numberOfOne)
		if1 := getOldPermutSum(input[1:], numberOfOne+1, numberOfPermut)
		return numberOfPermut + if0 + if1
	}
	panic("invalid character")
}
