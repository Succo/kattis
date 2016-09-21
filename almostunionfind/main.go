package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	list []int
	sets [][]int
	lens []int
	sums []int
	out  *bufio.Writer
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		s := &sets{make([]int, n), make([][]int, n), make([]int, n), make([]int, n), out}
		for i := range s.list {
			s.list[i] = i
			s.sets[i] = []int{i}
			s.lens[i] = 1
			s.sums[i] = i + 1
		}
		for i := 0; i < m; i++ {
			var com, p, q int
			fmt.Fscanln(in, &com, &p, &q)
			//fmt.Printf("%d %d %d\n", com, p, q)
			if com == 1 {
				s.union(p, q)
			} else if com == 2 {
				s.move(p, q)
			} else if com == 3 {
				s.output(p)
			}
			//fmt.Println(s.sets, s.list)
		}
		if in.Buffered() == 0 {
			break
		}
	}
	out.Flush()
}

func (s *sets) union(p, q int) {
	setp := s.list[p-1]
	setq := s.list[q-1]

	if setp == setq {
		return
	}

	for _, val := range s.sets[setp] {
		s.list[val] = setq
	}

	var set []int
	if s.lens[setp] > s.lens[setq] {
		set = append(s.sets[setp], s.sets[setq]...)
	} else {
		set = append(s.sets[setq], s.sets[setp]...)
	}
	s.sets[setq] = set

	s.lens[setq] += s.lens[setp]
	s.lens[setp] = 0

	s.sums[setq] += s.sums[setp]
	s.sums[setp] = 0
}

func (s *sets) move(p, q int) {
	setp := s.list[p-1]
	setq := s.list[q-1]

	if setp == setq {
		return
	}

	s.list[p-1] = setq
	s.sets[setq] = append(s.sets[setq], p-1)

	s.lens[setq]++
	s.lens[setp]--

	s.sums[setq] += p
	s.sums[setp] -= p

	if len(s.sets[setp]) == 1 {
		return
	}

	var pIndex int
	for i, val := range s.sets[setp] {
		if val == p-1 {
			pIndex = i
			break
		}
	}

	s.sets[setp][pIndex] = s.sets[setp][len(s.sets[setp])-1]
	s.sets[setp] = s.sets[setp][:len(s.sets[setp])-1]
}

func (s *sets) output(p int) {
	fmt.Fprintf(s.out, "%d %d\n", s.lens[s.list[p-1]], s.sums[s.list[p-1]])
}
