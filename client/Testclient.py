import unittest 
import client

class ClientSide(unittest.TestCase):   
    def test_Sudoku(self):
        N = 3
        sudokus = [client.generateSudoku() for k in range(0, N)]
        result = client.resolveSudoku([sudoku[1] for sudoku in sudokus])  
        for i in range(0, len(result)):
            if result[i] != None:
                for j in range(0, 9):
                    for k in range(0, 9):
                        if sudokus[i][1][j][k] != 0:
                            self.assertEqual(result[i][j][k], sudokus[i][1][j][k])

            for l in range(9):
                row = set()
                column = set()
                block = set()
                row_cube = 3 * (l//3)
                column_cube = 3 * (l%3)
                for j in range(9):
                    self.assertNotEqual(result[i][l][j], 0)
                    self.assertNotIn(result[i][l][j], row)
                    row.add(result[i][l][j])
                    self.assertNotIn(result[i][j][l], column)
                    column.add(result[i][j][l])
                    rc= row_cube+j//3
                    cc = column_cube + j%3
                    self.assertNotIn(result[i][rc][cc], block)
                    block.add(result[i][rc][cc])
if __name__ == '__main__': 
    unittest.main() 