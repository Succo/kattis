package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const (
	inputLength = 50
	numberOfIt  = 80
)

// TestBrutal test one versions of the function
// by brute forcing it using all combinaison possible
func TestBrutal(t *testing.T) {
	for i := 0; i < numberOfIt; i++ {
		iterate(t)
	}
}
func iterate(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	var buffer bytes.Buffer
	var k int
	var pos []int
	for i := 0; i < inputLength; i++ {
		char := chars[rand.Intn(length)]
		if char == '?' {
			pos = append(pos, i)
			k++
		}
		buffer.WriteByte(char)
	}
	buffer.WriteRune('\n')
	input := buffer.String()
	result, _, _ := getPermutSum(bufio.NewReader(strings.NewReader(input)))
	results := make([]int, power(2, k))
	for i := range results {
		results[i], _, _ = getPermutSum(bufio.NewReader(strings.NewReader(getInput(input, pos, i))))
	}
	expected := sum(results)
	if expected != result {
		t.Error(input)
	}
}

func getInput(input string, pos []int, i int) string {
	toMerge := fmt.Sprintf("%b", i)
	toMerge = leftPad(toMerge, len(pos))
	for j := range pos {
		input = input[:pos[j]] + string(toMerge[j]) + input[pos[j]+1:len(input)]
	}
	return input
}

func leftPad(input string, size int) string {
	toPad := size - len(input)
	for j := 0; j < toPad; j++ {
		input = "0" + input
	}
	return input
}

func power(x, y int) int {
	power := 1
	for i := 0; i < y; i++ {
		power = power * x
	}
	return power
}

func sum(in []int) int {
	var out int
	for _, val := range in {
		out += val
	}
	return out
}
