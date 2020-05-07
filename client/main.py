import client

def main():
    N = 2
    result = client.resolveSudoku([client.generateSudoku() for k in range(0, N)])
    for sudoku in result:
        for row in sudoku:
            print(row)
        print('\n')
    
main()