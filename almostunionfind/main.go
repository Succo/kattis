package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	list [][]int
	out  *bufio.Writer
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		s := &sets{make([][]int, n), out}
		for i := range s.list {
			s.list[i] = []int{i + 1}
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
	posp := -1
	posq := -1
	setp := -1
	setq := -1
	for i, set := range s.list {
		posi := pos(set, p)
		if posi > -1 {
			posp = posi
			setp = i
		}

		posi = pos(set, q)
		if posi > -1 {
			posq = posi
			setq = i
		}

		if posq > -1 && posp > -1 {
			break
		}
	}
	if setq == setp {
		return
	}
	s.list[setq] = append(s.list[setq], s.list[setp]...)
	s.list[setp] = []int{}
}

func (s *sets) move(p, q int) {
	posp := -1
	posq := -1
	setp := -1
	setq := -1
	for i, set := range s.list {
		posi := pos(set, p)
		if posi > -1 {
			posp = posi
			setp = i
		}

		posi = pos(set, q)
		if posi > -1 {
			posq = posi
			setq = i
		}

		if posq > -1 && posp > -1 {
			break
		}
	}
	if setq == setp {
		return
	}
	s.list[setp][posp] = s.list[setp][len(s.list[setp])-1]
	s.list[setp] = s.list[setp][:len(s.list[setp])-1]
	s.list[setq] = append(s.list[setq], p)

}

func (s *sets) output(p int) {
	for _, set := range s.list {
		posi := pos(set, p)
		if posi > -1 {
			fmt.Fprintf(s.out, "%d %d\n", len(set), sum(set))
			s.out.Flush()
			break
		}
	}
}

func sum(l []int) int {
	var res int
	for _, val := range l {
		res += val
	}
	return res
}

func pos(l []int, p int) int {
	for i, val := range l {
		if val == p {
			return i
		}
	}
	return -1
}
