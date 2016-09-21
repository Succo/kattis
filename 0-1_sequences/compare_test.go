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
	chars  = "01?"
	length = len(chars)
)

// TestCompare compare two versions of the function
// and print input that fails one of them
func TestCompare(t *testing.T) {
	// rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		var now time.Time
		len := rand.Intn(100)
		var buffer bytes.Buffer
		for j := 0; j < len; j++ {
			buffer.WriteByte(chars[rand.Intn(length)])
		}
		input := buffer.String()
		fmt.Println(input)
		now = time.Now()
		old := getOldPermutSum(input, 0, 0)
		fmt.Printf("old %v\n", time.Since(now))
		input = input + "\n"
		now = time.Now()
		new, _, _ := getPermutSum(bufio.NewReader(strings.NewReader(input)))
		fmt.Printf("new %v\n", time.Since(now))
		if old != new {
			fmt.Println("bug in the matrice")
			fmt.Println(input)
		}
	}
}
