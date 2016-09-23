package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	size  = 100000
	query = 100000
	cases = 100
)

func BenchmarkRandom(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var prints, unions, moves, init time.Duration
	file, _ := os.Open(os.DevNull)
	out := bufio.NewWriter(file)
	for c := 0; c < cases; c++ {
		now := time.Now()
		s := &sets{make([]int, size), make([][]int, size), make([]int, size), make([]int, size), out}
		for i := range s.dads {
			s.dads[i] = -1
			s.lens[i] = 1
			s.sons[i] = []int{}
			s.sums[i] = i + 1
		}
		init += time.Since(now)
		for i := 0; i < query; i++ {
			var com, p, q int
			com = rand.Intn(3) + 1
			p = rand.Intn(size) + 1
			q = rand.Intn(size) + 1
			if com == 1 {
				now := time.Now()
				s.union(p, q)
				unions += time.Since(now)
			} else if com == 2 {
				now := time.Now()
				s.move(p, q)
				moves += time.Since(now)
			} else if com == 3 {
				now := time.Now()
				s.output(p)
				prints += time.Since(now)
			}
		}
	}
	out.Flush()
	fmt.Printf("init %s\n", init.String())
	fmt.Printf("unions %s\n", unions.String())
	fmt.Printf("moves %s\n", moves.String())
	fmt.Printf("prints %s\n", prints.String())
}
