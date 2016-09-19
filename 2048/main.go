package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	board := make([][]int, 4)
	for i := 0; i < 4; i++ {
		var x0, x1, x2, x3 int
		fmt.Fscanln(in, &x0, &x1, &x2, &x3)
		board[i] = make([]int, 4)
		board[i][0] = x0
		board[i][1] = x1
		board[i][2] = x2
		board[i][3] = x3
	}
	var dir int
	fmt.Fscanln(in, &dir)
	if dir == 0 {
		for i := 0; i < 4; i++ {
			board[i][0], board[i][1], board[i][2], board[i][3] = compress(board[i][0], board[i][1], board[i][2], board[i][3])
		}
	} else if dir == 1 {
		for i := 0; i < 4; i++ {
			board[0][i], board[1][i], board[2][i], board[3][i] = compress(board[0][i], board[1][i], board[2][i], board[3][i])
		}
	} else if dir == 2 {
		for i := 0; i < 4; i++ {
			board[i][3], board[i][2], board[i][1], board[i][0] = compress(board[i][3], board[i][2], board[i][1], board[i][0])
		}
	} else if dir == 3 {
		for i := 0; i < 4; i++ {
			board[3][i], board[2][i], board[1][i], board[0][i] = compress(board[3][i], board[2][i], board[1][i], board[0][i])
		}
	}
	format(board)
}

func compress(x0, x1, x2, x3 int) (int, int, int, int) {
	merged := -1
	list := make([]int, 4)
	list[0] = x0
	list[1] = x1
	list[2] = x2
	list[3] = x3
	for i := 0; i < 4; i++ {
		val := list[i]
		if val == 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if list[j] == val && j > merged {
				list[j] = 2 * val
				list[j+1] = 0
				merged = j
			} else if list[j] == 0 {
				list[j] = val
				list[j+1] = 0
			} else {
				break
			}

		}
	}
	return list[0], list[1], list[2], list[3]
}

func format(board [][]int) {
	for _, line := range board {
		for _, val := range line {
			fmt.Print(val)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func isEmpty(list []int) bool {
	for _, val := range list {
		if val != 0 {
			return false
		}
	}
	return true
}
