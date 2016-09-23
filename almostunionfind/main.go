package main

import (
	"bufio"
	"fmt"
	"os"
)

type sets struct {
	dads []int
	sons [][]int
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
		for i := range s.dads {
			s.dads[i] = -1
			s.lens[i] = 1
			s.sons[i] = []int{}
			s.sums[i] = i + 1
		}
		for i := 0; i < m; i++ {
			var com, p, q int
			fmt.Fscanln(in, &com, &p, &q)
			//fmt.Printf("cmd %d %d %d\n", com, p, q)
			if com == 1 {
				s.union(p, q)
			} else if com == 2 {
				s.move(p, q)
			} else if com == 3 {
				s.output(p)
			}
			//s.drawTree()
			//fmt.Println(s.dads)
			//fmt.Println(s.sons)
			//s.assertTree()
		}
		if in.Buffered() == 0 {
			break
		}
	}
	out.Flush()
}

func (s *sets) getSet(p int) int {
	for s.dads[p] != -1 {
		p = s.dads[p]
	}
	return p
}

func (s *sets) union(p, q int) {
	setp := s.getSet(p - 1)
	setq := s.getSet(q - 1)

	if setp == setq {
		return
	}

	s.dads[setp] = setq
	s.sons[setq] = append(s.sons[setq], setp)

	s.lens[setq] += s.lens[setp]
	s.lens[setp] = 0

	s.sums[setq] += s.sums[setp]
	s.sums[setp] = 0
}

func (s *sets) move(p, q int) {
	setp := s.getSet(p - 1)
	setq := s.getSet(q - 1)
	if setp == setq {
		return
	}

	// i.e case were p "father" is p itself
	if setp == p-1 {
		// No sons, straightforward, assign to setq
		if len(s.sons[p-1]) == 0 {
			s.dads[setp] = setq
			s.sons[setq] = append(s.sons[setq], setp)

			s.lens[setq] += 1
			s.lens[setp] = 0

			s.sums[setq] += p
			s.sums[setp] = 0
		} else { // Need to rebuild a set around one of p sons
			root := s.sons[p-1][0]

			s.dads[root] = -1
			s.sons[root] = append(s.sons[root], s.sons[p-1][1:]...)

			for _, son := range s.sons[p-1][1:] {
				s.dads[son] = root
			}
			s.sons[p-1] = []int{}

			s.lens[root] = s.lens[p-1] - 1
			s.lens[p-1] = 0

			s.sums[root] = s.sums[p-1] - p
			s.sums[p-1] = 0

			s.dads[p-1] = setq
			s.lens[setq] += 1
			s.sums[setq] += p
			s.sons[setq] = append(s.sons[setq], p-1)
		}
	} else {
		root := s.dads[p-1]
		sons := s.sons[root]
		s.lens[setp] -= 1
		s.sums[setp] -= p
		pIndex := -1
		for i, son := range s.sons[root] {
			if son == p-1 {
				pIndex = i
			}
		}

		if pIndex == -1 {
			s.drawTree()
			fmt.Println(sons)
			fmt.Println(p)
			fmt.Println(root)
			panic("should not happen")
		}
		s.dads[p-1] = setq
		s.lens[setq] += 1
		s.sums[setq] += p
		s.sons[setq] = append(s.sons[setq], p-1)

		if len(s.sons[p-1]) == 0 {
			sons[pIndex] = sons[len(sons)-1]
			s.sons[root] = sons[:len(sons)-1]
		} else {
			sons[pIndex] = s.sons[p-1][0]
			s.sons[root] = append(sons, s.sons[p-1][1:]...)
			for _, son := range s.sons[p-1] {
				s.dads[son] = root
			}
			s.sons[p-1] = []int{}
		}
	}
}

func (s *sets) output(p int) {
	set := s.getSet(p - 1)
	fmt.Fprintf(s.out, "%d %d\n", s.lens[set], s.sums[set])
}

func (s *sets) totalSum() int {
	var total int
	for _, sum := range s.sums {
		total += sum
	}
	return total
}

func (s *sets) totalLen() int {
	var total int
	for _, len := range s.lens {
		total += len
	}
	return total
}

func (s *sets) drawTree() {
	for i, sum := range s.sums {
		if sum != 0 {
			s.draw(i)
			fmt.Println("")
		}
	}
}

func (s *sets) draw(i int) {
	fmt.Printf("%d ", i)
	for _, son := range s.sons[i] {
		s.draw(son)
	}
}

func (s *sets) assertTree() {
	for i, dad := range s.dads {
		if dad != -1 {
			s.assertFils(dad, i)
		}
	}
}

func (s *sets) assertFils(dad, i int) {
	for _, fil := range s.sons[dad] {
		if fil == i {
			return
		}
	}
	panic(fmt.Sprintf("improper fils %d dad %d", i, dad))
}
