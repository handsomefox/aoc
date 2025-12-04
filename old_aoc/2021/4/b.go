package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Board2 struct {
	rows [BOARD_SIZE]Row
	won  bool
}

type Winner2 struct {
	board     *Board2
	index     int
	finalDraw int
}

func NewBoard2(rows [BOARD_SIZE]Row) *Board2 {
	return &Board2{
		rows: rows,
		won:  false,
	}
}

func SolveB(input string) string {
	sn := bufio.NewScanner(strings.NewReader(input))

	order := MakeDraws(sn)
	// fmt.Println(order)

	boards := MakeBoards2(sn)
	// PrettyPrintBoards(boards)

	winner := DrawValues2(order, boards)
	sumOfAllUnmarked := FindSumOfAllUnmarked2(winner)
	return fmt.Sprintf("Result: %v", sumOfAllUnmarked*winner.finalDraw)
}

func makeBoard2(sn *bufio.Scanner) *Board2 {
	var rows [BOARD_SIZE]Row
	for i := 0; i < BOARD_SIZE; i++ {
		sn.Scan()
		splt := strings.Fields(sn.Text())

		numbers := make([]int, 0, 5)
		for _, v := range splt {
			i, _ := strconv.Atoi(v)
			numbers = append(numbers, i)
		}

		row := NewRow(numbers)
		rows[i] = row
	}

	return NewBoard2(rows)
}

func MakeBoards2(sn *bufio.Scanner) []Board2 {
	boards := make([]Board2, 0)

	for sn.Scan() {
		newBoard := makeBoard2(sn)
		boards = append(boards, *newBoard)
	}
	return boards
}

func DrawValues2(draws []int, boards []Board2) *Winner2 {
	var winner *Winner2
	for _, draw := range draws {
		markDrawnValueOnBoards2(draw, boards)
		winner = tryFindWinnerBoard2(draw, boards)
		if winner != nil {
			if allBoardsWon2(boards) {
				return winner
			}
		}
	}
	return winner
}

func allBoardsWon2(boards []Board2) bool {
	for _, v := range boards {
		if !v.won {
			return false
		}
	}
	return true
}

func markDrawnValueOnBoards2(drawn int, boards []Board2) {
	for boardIndex, board := range boards {
		for rowIndex, row := range board.rows {
			for valueIndex, values := range row.nums {
				if values.Val == drawn {
					boards[boardIndex].rows[rowIndex].nums[valueIndex].Marked = true
				}
			}
		}
	}
}

func tryFindWinnerBoard2(drawn int, boards []Board2) *Winner2 {
	var winner *Winner2
	for i, board := range boards {
		if board.won {
			continue
		}
		if checkColumns2(board) || checkRows2(board) {
			boards[i].won = true
			winner = &Winner2{
				board:     &boards[i],
				index:     i,
				finalDraw: drawn,
			}
		}
	}
	return winner
}

func checkColumns2(board Board2) bool {
	// Check every column
	for i := 0; i < BOARD_SIZE; i++ {
		won := true
		for j := 0; j < BOARD_SIZE; j++ {
			val := board.rows[j].nums[i]
			// If at least one value is unmarked, this column can't win.
			if !val.Marked {
				won = false
			}
		}
		// If all the values were marked, this column won, we can say that this board has won.
		if won {
			return won
		}
	}
	return false
}

func checkRows2(board Board2) bool {
	// Check every row
	for _, row := range board.rows {
		won := true
		// If at least one value is unmarked, this row can't win.
		for _, values := range row.nums {
			if !values.Marked {
				won = false
			}
		}
		// If all the values were marked, this row won, we can say that this board has won.
		if won {
			return won
		}
	}
	return false
}

func FindSumOfAllUnmarked2(winner *Winner2) int {
	var sumOfAllUnmarked int
	for _, r := range winner.board.rows {
		for _, v := range r.nums {
			if !v.Marked {
				sumOfAllUnmarked += v.Val
			}
		}
	}
	return sumOfAllUnmarked
}
