impl Solution {
    pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
        let mut res: Vec<Vec<String>> = Vec::new();
        let n = n as usize;
        let mut board: Vec<Vec<char>> = Vec::new();
        for _ in 0..n {
            board.push(vec!['.'; n as usize]);
        }

        Solution::helper(&mut board, 0, n, &mut res);

        res
    }

    fn helper(board: &mut Vec<Vec<char>>, ind: usize, n: usize, res: &mut Vec<Vec<String>>) {
        if ind == n {
            Solution::add_board(board.clone(), res);
        }

        for i in 0..n {
            if Solution::is_valid(ind, i, board, n) {
                board[ind][i] = 'Q';
                Solution::helper(board, ind + 1, n, res);
                board[ind][i] = '.';
            }
        }
    }

    fn add_board(board: Vec<Vec<char>>, res: &mut Vec<Vec<String>>) {
        let mut new_board: Vec<String> = Vec::new();

        for item in board {
            new_board.push(item.iter().collect());
        }

        res.push(new_board);
    }

    fn is_valid(row: usize, col: usize, board: &mut Vec<Vec<char>>, n: usize) -> bool {
        if row == 0 {
            return true;
        }

        //up
        let mut r = row - 1;
        while r >= 0 {
            if board[r][col] == 'Q' {
                return false;
            }
            if r == 0 {
                break;
            }
            r -= 1
        }

        //right diagonal
        let (mut r, mut c) = (row - 1, col + 1);
        while r >= 0 && c < n {
            if board[r][c] == 'Q' {
                return false;
            }
            if r == 0 {
                break;
            }
            r -= 1;
            c += 1;
        }

        if col == 0 {
            return true;
        }

        //left diagonal
        let (mut r, mut c) = (row - 1, col - 1);
        while r >= 0 && c >= 0 {
            if board[r][c] == 'Q' {
                return false;
            }
            if r == 0 || c == 0 {
                break;
            }
            r -= 1;
            c -= 1;
        }

        true
    }
}
