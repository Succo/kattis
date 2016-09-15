package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	mod = 1000000007
)

func getPermutSum(input string) (int, int, int) {
	if len(input) == 0 {
		return 0, 0, 1
	}
	if input[0] == '0' {
		permuts, zeros, paths := getPermutSum(input[1:])
		return permuts, zeros + paths, paths
	}
	if input[0] == '1' {
		permut, zeros, paths := getPermutSum(input[1:])
		return permut + zeros, zeros, paths
	}
	if input[0] == '?' {
		permut, zeros, paths := getPermutSum(input[1:])
		return 2*permut + zeros, 2*zeros + paths, paths * 2
	}
	panic("invalid character")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	now := time.Now()

	input, err := in.ReadString('\n')
	if err != nil {
		panic("error reading stdin")
	}

	input = input[:len(input)-1]
	sum, _, _ := getPermutSum(input)
	fmt.Printf("%d\n", sum%mod)
	fmt.Println(time.Since(now))
}
