import client

def main():
    sudoku = [ 
                [[8, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 0, 3, 6, 0, 0, 0, 0, 0],
                [0, 7, 0, 0, 9, 0, 2, 0, 0],
                [0, 5, 0, 0, 0, 7, 0, 0, 0],
                [0, 0, 0, 0, 4, 5, 7, 0, 0],
                [0, 0, 0, 1, 0, 0, 0, 3, 0],
                [0, 0, 1, 0, 0, 0, 0, 6, 8],
                [0, 0, 8, 5, 0, 0, 0, 1, 0],
                [0, 9, 0, 0, 0, 0, 4, 0, 0]],
                [
                [6, 0, 0, 0, 5, 0, 0, 7, 8],
                [0, 9, 0, 0, 0, 1, 5, 0, 0],
                [0, 7, 0, 0, 9, 0, 3, 0, 1],
                [0, 0, 0, 0, 0, 5, 0, 0, 0],
                [0, 0, 0, 8, 1, 3, 0, 0, 0],
                [0, 0, 0, 6, 0, 0, 0, 0, 0],
                [5, 0, 7, 0, 8, 0, 0, 2, 0],
                [0, 0, 9, 4, 0, 0, 0, 1, 0],
                [1, 4, 0, 0, 3, 0, 0, 0, 6]]]

    result = client.resolveSudoku(sudoku)
    for sudokuNotSolvedYet, sudokuGet in zip(sudoku, result):
        if sudokuGet != None:
            print("Sudoku : ")
            for row in sudokuNotSolvedYet:
                print(row)
            print("Solution : ")
            for row in sudokuGet:
                print(row)
        else:
            print("Sudoku : ")
            for row in sudokuNotSolvedYet:
                print(row)
            print("No solution!")
        print('\n')
    
main()