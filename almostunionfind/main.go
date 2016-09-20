package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	list []map[int]bool
	out  *bufio.Writer
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		s := &sets{make([]map[int]bool, n), out}
		for i := range s.list {
			s.list[i] = map[int]bool{i + 1: true}
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
	setp := -1
	setq := -1
	for key, set := range s.list {
		_, ok := set[p]
		if ok {
			setp = key
		}
		_, ok = set[q]
		if ok {
			setq = key
		}
		if setp != -1 && setq != -1 {
			break
		}
	}
	for key, _ := range s.list[setp] {
		s.list[setq][key] = true
	}
	s.list[setp] = make(map[int]bool)
}

func (s *sets) move(p, q int) {
	setp := -1
	setq := -1
	for key, set := range s.list {
		_, ok := set[p]
		if ok {
			setp = key
		}
		_, ok = set[q]
		if ok {
			setq = key
		}
		if setp != -1 && setq != -1 {
			break
		}
	}
	s.list[setq][p] = true
	delete(s.list[setp], p)
}

func (s *sets) output(p int) {
	for _, set := range s.list {
		_, ok := set[p]
		if ok {
			fmt.Fprintf(s.out, "%d %d\n", len(set), sum(set))
			s.out.Flush()
			break
		}
	}
}

func sum(set map[int]bool) int {
	var sum int
	for key, _ := range set {
		sum += key
	}
	return sum
}
