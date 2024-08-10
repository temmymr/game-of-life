package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomState(width int, height int) [][]int {
	board := make([][]int, width)
	for i := 0; i < width; i++ {
		board[i] = make([]int, height)
		for j := 0; j < height; j++ {
			board[i][j] = rand.Intn(2)
		}
	}
	return board
}

func render(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			switch board[i][j] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
}

func countNeighbor(board [][]int, x, y int) int {
	count := 0

	var dir = [][]int{
		{-1, -1}, // LU
		{0, -1},  // MU
		{1, -1},  // RU
		{-1, 0},  // LM
		{1, 0},   // RM
		{-1, 1},  // LD
		{0, 1},   // MD
		{1, 1},   // RD
	}

	for _, d := range dir {
		column := x + d[0]
		row := y + d[1]

		if column < 0 || row < 0 {
			continue
		}

		if column >= len(board[y]) || row >= len(board) {
			continue
		}

		if board[row][column] == 1 {
			count++
		}
	}
	return count
}

func nextBoardState(board [][]int) {
	nextState := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		nextState[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			nextState[i][j] = board[i][j]
			neighbor := countNeighbor(board, j, i)

			if neighbor < 2 {
				nextState[i][j] = 0
			}

			if neighbor > 3 {
				nextState[i][j] = 0
			}

			if neighbor == 3 {
				nextState[i][j] = 1
			}
		}
	}

	copy(board, nextState)
}

func main() {
	board := randomState(24, 80)
	ticker := time.NewTicker(80 * time.Millisecond)

	for {
		<-ticker.C
		fmt.Printf("\x1bc") // clear console

		render(board)
		nextBoardState(board)
	}
}
