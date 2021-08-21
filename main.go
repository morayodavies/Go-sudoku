package main

import (
	"fmt"
	"os"
)

func main() {

	board := parseInput(os.Args[1])

	if backtrack(&board) {
		printBoard(board)
	} else {
		fmt.Printf("Error")
	}
}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
			}
			fmt.Printf("%d ", board[row][col])
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("")
		} else {
			fmt.Println()
		}
	}
}

func createboard(sudokustring string) [9]int {
	var answer [9]int
	for i, letter := range sudokustring {
		if letter == '.' {
			answer[i] = 0
		} else {
			answer[i] = int(letter) - 48
		}
	}
	return answer
}

func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	sudokulines := os.Args[1:]
	for i := 0; i < 9; i++ {
		board[i] = createboard(sudokulines[i])
	}
	return board
}
