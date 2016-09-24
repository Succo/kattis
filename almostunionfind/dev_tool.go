package main

import "fmt"

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
