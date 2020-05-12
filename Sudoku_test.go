package main

import (
	"reflect"
	"testing"
)

var boardEasySudoku = &Board{
	Board: [9][9]int{
		{1, 0, 3, 0, 0, 6, 0, 8, 0},
		{0, 5, 0, 0, 8, 0, 1, 2, 0},
		{7, 0, 9, 1, 0, 3, 0, 5, 6},
		{0, 3, 0, 0, 6, 7, 0, 9, 0},
		{5, 0, 7, 8, 0, 0, 0, 3, 0},
		{8, 0, 1, 0, 3, 0, 5, 0, 7},
		{0, 4, 0, 0, 7, 8, 0, 1, 0},
		{6, 0, 8, 0, 0, 2, 0, 4, 0},
		{0, 1, 2, 0, 4, 5, 0, 7, 8},
	},
	Success: false,
}

var solutionEasySudoku = &Board{
	Board: [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	},
	Success: false,
}

var boardHardSudoku = &Board{
	Board: [9][9]int{
		{0, 6, 1, 0, 0, 7, 0, 0, 3},
		{0, 9, 2, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 5, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 4},
		{5, 0, 0, 0, 0, 8, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 6, 0, 8, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	Success: false,
}

var solutionHardSudoku = &Board{
	Board: [9][9]int{
		{4, 6, 1, 9, 8, 7, 2, 5, 3},
		{7, 9, 2, 4, 5, 3, 1, 6, 8},
		{3, 8, 5, 2, 1, 6, 4, 7, 9},
		{1, 2, 8, 5, 3, 4, 7, 9, 6},
		{9, 3, 6, 7, 2, 1, 5, 8, 4},
		{5, 7, 4, 6, 9, 8, 3, 1, 2},
		{8, 4, 9, 3, 7, 5, 6, 2, 1},
		{2, 5, 3, 1, 6, 9, 8, 4, 7},
		{6, 1, 7, 8, 4, 2, 9, 3, 5},
	},
	Success: false,
}

var boardNotPossible = &Board{
	Board: [9][9]int{
		{0, 6, 1, 0, 0, 7, 0, 0, 3},
		{0, 9, 2, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 5, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 4},
		{5, 0, 0, 0, 0, 8, 0, 5, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 6, 0, 8, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	Success: false,
}

func copyBoard(cells [9][9]int) *Board {
	return &Board{Board: cells, Success: false}
}

func TestBoard_Backtrack(t *testing.T) {
	for i, check := range []struct {
		board, expect *Board
	}{
		{boardEasySudoku, solutionEasySudoku},
		{boardHardSudoku, solutionHardSudoku},
	} {
		// create copy the board since Backtrack() modifies the orignal values
		copy := copyBoard(check.board.Board)
		if !copy.solve() || !reflect.DeepEqual(copy.Board, check.expect.Board) {
			t.Errorf("[%d] the given board wasn't solved as expected", i)
		}
	}
}

func TestBoard_isPosibleToSolve(t *testing.T) {
	for i, check := range []struct {
		board    *Board
		expected bool
	}{
		{boardHardSudoku, true},
		{boardNotPossible, false},
	} {
		result := check.board.isPosibleToSolve()
		if result != check.expected {
			t.Errorf("[%d] expect to find empty cell", i)
		}
	}
}

func BenchmarkBacktrackEasySudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// prevent the modification of the orignal board
		copy := copyBoard(boardEasySudoku.Board)
		b.StartTimer()
		copy.solve()
	}
}

func BenchmarkBacktrackHardSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// prevent the modification of the orignal board
		copy := copyBoard(boardHardSudoku.Board)
		b.StartTimer()
		copy.solve()
	}
}
