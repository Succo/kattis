package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	list map[int]int
	out  *bufio.Writer
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		s := &sets{make(map[int]int), out}
		for i := 0; i < n; i++ {
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
	set := s.list[p-1]
	for key, value := range s.list {
		if value == set {
			s.list[key] = s.list[q-1]
		}
	}
}

func (s *sets) move(p, q int) {
	s.list[p-1] = s.list[q-1]
}

func (s *sets) output(p int) {
	var sum, len int
	for key, value := range s.list {
		if value == s.list[p-1] {
			sum += key + 1
			len++
		}
	}
	fmt.Fprintf(s.out, "%d %d\n", len, sum)
	s.out.Flush()
}
