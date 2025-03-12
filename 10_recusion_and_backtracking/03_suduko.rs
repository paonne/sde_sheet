impl Solution {
    pub fn solve_sudoku(board: &mut Vec<Vec<char>>) {
        Solution::helper(board);
    }

    fn helper(b: &mut Vec<Vec<char>>) -> bool {
        for i in 0..9 {
            for j in 0..9 {
                if b[i][j] == '.' {
                    for c in "123456789".chars() {
                        if Solution::is_valid(i, j, b, c) {
                            b[i][j] = c;
                            if Solution::helper(b) == true {
                                return true;
                            }
                            b[i][j] = '.';
                        }
                    }
                    return false;
                }
            }
        }

        true
    }

    fn is_valid(r: usize, c: usize, b: &Vec<Vec<char>>, digit: char) -> bool {
        for i in 0..9 {
            for j in 0..9 {
                if b[r][j] == digit || b[i][c] == digit {
                    return false;
                }
            }
        }

        let (grid_row, grid_col) = ((r / 3) * 3, (c / 3) * 3);
        for i in grid_row..grid_row + 3 {
            for j in grid_col..grid_col + 3 {
                if b[i][j] == digit {
                    return false;
                }
            }
        }

        true
    }
}
