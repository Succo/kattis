package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := 1000
	c := 1000
	n := 1000
	fmt.Printf("%d %d\n", r, c)
	for i := 0; i < r; i++ {
		var line bytes.Buffer
		for j := 0; j < c; j++ {
			line.WriteString(strconv.Itoa((i + j) % 2))
		}
		fmt.Println(line.String())
	}
	fmt.Printf("%d\n", n)
	for i := 0; i < n; i++ {
		var line bytes.Buffer
		line.WriteString(strconv.Itoa(rand.Intn(r) + 1))
		line.WriteString(" ")
		line.WriteString(strconv.Itoa(rand.Intn(c) + 1))
		line.WriteString(" ")
		line.WriteString(strconv.Itoa(rand.Intn(r) + 1))
		line.WriteString(" ")
		line.WriteString(strconv.Itoa(rand.Intn(c) + 1))
		fmt.Println(line.String())
	}
}
