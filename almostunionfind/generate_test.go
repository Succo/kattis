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

func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var prints, unions, moves time.Duration
	file, _ := os.Open(os.DevNull)
	out := bufio.NewWriter(file)
	for c := 0; c < cases; c++ {
		s := &sets{make([]int, size), make([][]int, size), make([]int, size), make([]int, size), out}
		for i := range s.list {
			s.list[i] = i
			s.sets[i] = []int{i}
		}
		for i := 0; i < query; i++ {
			var com, p, q int
			com = rand.Intn(3) + 1
			//com = 1
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
	fmt.Printf("unions %s\n", unions.String())
	fmt.Printf("moves %s\n", moves.String())
	fmt.Printf("prints %s\n", prints.String())
	fmt.Printf("list update %s\n", listUpdate.String())
	fmt.Printf("set update %s\n", setUpdate.String())
	fmt.Printf("pRemoval %s\n", pRemoval.String())
}
