package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	list []int
	sets [][]int
	out  *bufio.Writer
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		s := &sets{make([]int, n), make([][]int, n), out}
		for i := range s.list {
			s.list[i] = i
			s.sets[i] = []int{i}
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

	set := make([]int, len(s.sets[setq])+len(s.sets[setp]))
	copy(set, s.sets[setq])
	copy(set[len(s.sets[setq]):], s.sets[setp])
	s.sets[setq] = set

	s.sets[setp] = make([]int, 0)
}

func (s *sets) move(p, q int) {
	setp := s.list[p-1]
	setq := s.list[q-1]

	if setp == setq {
		return
	}

	s.list[p-1] = setq
	s.sets[setq] = append(s.sets[setq], p-1)

	if len(s.sets[setp]) == 1 {
		s.sets[setp] = make([]int, 0)
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
	var sum, len int
	setp := s.list[p-1]
	for _, val := range s.sets[setp] {
		sum += val + 1
		len++
	}
	fmt.Fprintf(s.out, "%d %d\n", len, sum)
	s.out.Flush()
}
