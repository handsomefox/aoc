package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

type Value struct {
	Val    int
	Marked bool
}

type Row struct {
	nums [BOARD_SIZE]Value
}

func NewRow(nums []int) Row {
	row := Row{}
	for i := 0; i < BOARD_SIZE; i++ {
		row.nums[i] = Value{
			Val:    nums[i],
			Marked: false,
		}
	}
	return row
}

type Board struct {
	rows [BOARD_SIZE]Row
}

type Winner struct {
	board     *Board
	index     int
	finalDraw int
}

func NewBoard(rows [BOARD_SIZE]Row) *Board {
	return &Board{
		rows: rows,
	}
}

func SolveA(input string) string {
	sn := bufio.NewScanner(strings.NewReader(input))

	var (
		order  = MakeDraws(sn)
		boards = MakeBoards(sn)
		winner = DrawValues(order, boards)
	)

	sumOfAllUnmarked := FindSumOfAllUnmarked(winner)
	return fmt.Sprintf("Result: %v", sumOfAllUnmarked*winner.finalDraw)
}

func MakeDraws(sn *bufio.Scanner) []int {
	sn.Scan()
	orderStr := strings.Split(sn.Text(), ",")

	ret := make([]int, 0, cap(orderStr))

	for _, v := range orderStr {
		i, _ := strconv.Atoi(v)
		ret = append(ret, i)
	}
	return ret
}

func makeBoard(sn *bufio.Scanner) *Board {
	var rows [BOARD_SIZE]Row
	for i := 0; i < BOARD_SIZE; i++ {
		sn.Scan()
		split := strings.Fields(sn.Text())

		numbers := make([]int, 0, 5)
		for _, v := range split {
			i, _ := strconv.Atoi(v)
			numbers = append(numbers, i)
		}

		row := NewRow(numbers)
		rows[i] = row
	}

	return NewBoard(rows)
}

func MakeBoards(sn *bufio.Scanner) []Board {
	boards := make([]Board, 0)

	for sn.Scan() {
		newBoard := makeBoard(sn)
		boards = append(boards, *newBoard)
	}
	return boards
}

func PrettyPrintRow(row *Row) {
	for i := 0; i < BOARD_SIZE; i++ {
		if row.nums[i].Marked {
			fmt.Printf("\033[1m%v\033[0m ", row.nums[i].Val)
		} else {
			fmt.Printf("%v ", row.nums[i].Val)
		}
	}
	fmt.Println()
}

func PrettyPrintBoards(boards []Board) {
	for i, board := range boards {
		fmt.Printf("Board #%v\n", i+1)
		for _, r := range board.rows {
			PrettyPrintRow(&r)
		}
	}
}

func PrettyPrintWinner(winner *Winner) {
	fmt.Printf("Board #%v has won!\n", winner.index+1)
	for _, r := range winner.board.rows {
		PrettyPrintRow(&r)
	}
}

func DrawValues(draws []int, boards []Board) *Winner {
	var winner *Winner
	for _, draw := range draws {
		markDrawnValueOnBoards(draw, boards)
		winner = tryFindWinnerBoard(draw, boards)
		if winner != nil {
			return winner
		}
	}
	return winner
}

func markDrawnValueOnBoards(drawn int, boards []Board) {
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

func tryFindWinnerBoard(drawn int, boards []Board) *Winner {
	for i, board := range boards {
		if checkColumns(board) || checkRows(board) {
			return &Winner{
				board:     &boards[i],
				index:     i,
				finalDraw: drawn,
			}
		}
	}
	return nil
}

func checkColumns(board Board) bool {
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

func checkRows(board Board) bool {
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

func FindSumOfAllUnmarked(winner *Winner) int {
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
