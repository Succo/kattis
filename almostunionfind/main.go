package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	list []int
	out  *bufio.Writer
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		s := &sets{make([]int, n), out}
		for i := range s.list {
			s.list[i] = i
		}
		for i := 0; i < m; i++ {
			var com, p, q int
			fmt.Fscanln(in, &com, &p, &q)
			if com == 1 {
				s.union(p, q)
			} else if com == 2 {
				s.move(p, q)
			} else if com == 3 {
				s.output(p)
			}
		}
		if in.Buffered() == 0 {
			break
		}
	}
}

func (s *sets) union(p, q int) {
	setp := s.list[p-1]
	setq := s.list[q-1]
	for key, set := range s.list {
		if set == setp {
			s.list[key] = setq
		}
	}
}

func (s *sets) move(p, q int) {
	s.list[p-1] = s.list[q-1]
}

func (s *sets) output(p int) {
	var sum, len int
	setp := s.list[p-1]
	for key, set := range s.list {
		if set == setp {
			sum += key + 1
			len++
		}
	}
	fmt.Fprintf(s.out, "%d %d\n", len, sum)
	s.out.Flush()
}
