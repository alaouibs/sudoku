import json
import requests

import copy
def generateSudoku():
    ''' 
    generate sudoku
    '''
    base = 3
    side = base*base

    # pattern for a baseline valid solution
    def pattern(r,c): return (base*(r%base)+r//base+c)%side

    # randomize rows, columns and numbers (of valid base pattern)
    from random import sample
    def shuffle(s): return sample(s,len(s)) 
    rBase = range(base) 
    rows  = [ g*base + r for g in shuffle(rBase) for r in shuffle(rBase) ] 
    cols  = [ g*base + c for g in shuffle(rBase) for c in shuffle(rBase) ]
    nums  = shuffle(range(1,base*base+1))

    # produce board using randomized baseline pattern
    board = [ [nums[pattern(r,c)] for c in cols] for r in rows ]
    excepted = copy.deepcopy(board)
    squares = side*side
    empties = squares * 3//4
    for p in sample(range(squares),empties):
        board[p//side][p%side] = 0
    return excepted, board

def resolveSudoku(sudokus):
    ''' 
    Solves a List or Sudoku by asking a Server
    '''
    r = requests.post('http://localhost:10000/sudoku', json=[{"Board": sudoku, "Success": False} for sudoku in sudokus])
    result = r.json()

    for i in range(0, len(result)):
        if result[i]['Success']:
            result[i] = result[i]['Board']
        else:
            result[i] = None
    return result
