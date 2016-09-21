package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	min = "abcdefghijklmnopqrstuvwxyz"
	maj = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 10; i++ {
		lines := 500 + rand.Intn(500)
		bug := randString(rand.Intn(90)+1, maj)
		fmt.Printf("%d %s\n", lines, bug)
		for j := 0; j < lines; j++ {
			fmt.Println(randLines(bug))
			if j%10 == 0 {
				fmt.Fprintf(os.Stderr, "line %d \n", j)
			}
		}
		fmt.Fprintf(os.Stderr, "test case %d \n", i)
	}
}

func randString(n int, letter string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func randLines(bug string) string {
	var buffer bytes.Buffer
	for i := 0; i < (rand.Intn(1000) + 1000); i++ {
		if rand.Intn(10) > 9 {
			buffer.WriteString(bug)
		} else {
			buffer.WriteByte(min[rand.Intn(len(min))])
		}
	}
	return buffer.String()
}
