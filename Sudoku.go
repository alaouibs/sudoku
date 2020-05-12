package main

type Board struct {
	Board   [9][9]int
	Success bool
}

func (board *Board) isNumberInRow(currRow int, currColumn int, posibleNumber int) bool {
	for column := 0; column < 9; column++ {
		if (column != currColumn) && (board.Board[currRow][column] == posibleNumber) {
			return true
		}
	}
	return false
}

func (board *Board) isNumberInColumn(currRow int, currColumn int, posibleNumber int) bool {
	for row := 0; row < 9; row++ {
		if (row != currRow) && (board.Board[row][currColumn] == posibleNumber) {
			return true
		}
	}
	return false
}

func (board *Board) isNumberInSubGrid(currRow int, currColumn int, posibleNumber int) bool {
	startIndexRowSubGrid := currRow - currRow%3
	endIndexRowSubGrid := startIndexRowSubGrid + 3

	startIndexColumnSubGrid := currColumn - currColumn%3
	endIndexColumnSubGrid := startIndexColumnSubGrid + 3
	for row := startIndexRowSubGrid; row < endIndexRowSubGrid; row++ {
		for column := startIndexColumnSubGrid; column < endIndexColumnSubGrid; column++ {
			if (column != currColumn) && (row != currRow) && (board.Board[row][column] == posibleNumber) {
				return true
			}
		}
	}
	return false
}

func (board *Board) isNumberPosible(row int, column int, posibleNumber int) bool {
	return !board.isNumberInRow(row, column, posibleNumber) &&
		!board.isNumberInColumn(row, column, posibleNumber) &&
		!board.isNumberInSubGrid(row, column, posibleNumber)
}

func (board *Board) isPosibleToSolve() bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			currCell := board.Board[row][column]
			if currCell != 0 {
				if !board.isNumberPosible(row, column, currCell) {
					return false
				}
			}
		}
	}
	return true
}

func (board *Board) solve() bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board.Board[row][column] == 0 {
				for posibleNumber := 1; posibleNumber <= 9; posibleNumber++ {
					if board.isNumberPosible(row, column, posibleNumber) {
						board.Board[row][column] = posibleNumber

						if !board.solve() {
							board.Board[row][column] = 0
						} else {
							board.Success = true
							return true
						}
					}
				}
				board.Success = false
				return false
			}
		}
	}
	board.Success = true
	return true
}
