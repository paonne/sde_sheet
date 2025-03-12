## Print All Permutations of a String/Array

**❓**: Print All Permutations of a String/Array  

**Example**:  

Input: arr = [1, 2, 3]

Output: 
[
  [1, 2, 3],
  [1, 3, 2],
  [2, 1, 3],
  [2, 3, 1],
  [3, 1, 2],
  [3, 2, 1]
]  

**🧠**:  
1. Keep track of seen elements.  
2. At every index, try all possible.
=== "🐍"
    ```py
    --8<-- "10_recusion_and_backtracking/01_permutations.py"
    ```
=== "🦀"
    ```rust
    --8<-- "10_recusion_and_backtracking/01_permutations.rs"
    ```
[📘](https://takeuforward.org/data-structure/print-all-permutations-of-a-string-array/) [💻](https://leetcode.com/problems/permutations/description/)<br>

---

## N Queen Problem

**❓**: The n-queens is the problem of placing n queens on n × n chessboard such that no two queens can attack each other. Given an integer n, return all distinct solutions to the n -queens puzzle. Each solution contains a distinct boards configuration of the queen's placement, where ‘Q’ and ‘.’ indicate queen and empty space respectively.

**Example**:  
Input: n = 4

Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
![NQueens](https://lh5.googleusercontent.com/Llq0kUywtAlP5K6dFU_DUSdsPSUhWK3zsnBpc8U85nqqykaKqJ1GPoiPIExFOkFCGDnQJj2zKEmlUMiAJlCwFCK1y9SVg4YMy4tg2P9EmbeG2uLcyHFmBKX_cW5eTJJ2dkEut4bx)

**🧠**:  
1. At every position, check whether it is a valid position or not.  
2. If yes, repeat process until we reach end index.  
3. If no, backtrack and repeat the process
=== "🐍"
    ```py
    --8<-- "10_recusion_and_backtracking/02_n_queens.py"
    ```
=== "🦀"
    ```rust
    --8<-- "10_recusion_and_backtracking/02_n_queens.rs"
    ```
[📘](https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/) [💻](https://leetcode.com/problems/n-queens/description/)<br>

---

## Sudoku Solver

**❓**: Given a 9x9 incomplete sudoku, solve it such that it becomes valid sudoku.   

**Example**:  
![suduko](https://lh5.googleusercontent.com/k78fKDRjAJU3CIBgMRYCDEG93ndte0k85JLWYK6IumRreKBRv5zcKDkc1Ms_E6Bi_2M4twPY5GWos_0kQNkZO9AXRtowc5sKe5KZMJpcCqKddtXDr7xuA-HZDIttJ_-5RE30NlDJ)
=== "🐍"
    ```py
    --8<-- "10_recusion_and_backtracking/03_suduko.py"
    ```
=== "🦀"
    ```rust
    --8<-- "10_recusion_and_backtracking/03_suduko.rs"
    ```
[📘](https://takeuforward.org/data-structure/sudoku-solver/) [💻](https://leetcode.com/problems/sudoku-solver/)<br>

---

## M Coloring Problem

**❓**: Given an undirected graph and a number m, determine if the graph can be colored with at most m colors such that no two adjacent vertices of the graph are colored with the same color.  

**Example**:  

Input: 
N = 4
M = 3
E = 5

Edges[] = {
  (0, 1),
  (1, 2),
  (2, 3),
  (3, 0),
  (0, 2)
}

Output: 1

**🧠**:  
1. Form the graph as adjacency list.  
2. For every node, try to color all the colors.  
3. If able to color the node with particular color, store it and move to next node.  
4. Once all nodes are colored, return True, else return False.  

```py
{!10_recusion_and_backtracking/04_m_coloring.py!}
```

[📘](https://takeuforward.org/data-structure/m-coloring-problem/)<br>

---

## Rat in a Maze

**❓**: Consider a rat placed at (0, 0) in a square matrix of order N * N. It has to reach the destination at (N - 1, N - 1). Find all possible paths that the rat can take to reach from source to destination. The directions in which the rat can move are 'U'(up), 'D'(down), 'L' (left), 'R' (right). Value 0 at a cell in the matrix represents that it is blocked and the rat cannot move to it while value 1 at a cell in the matrix represents that rat can travel through it.

Note: In a path, no cell can be visited more than one time.

Print the answer in lexicographical(sorted) order

**Example**:  
Input:
N = 4
m[][] = {{1, 0, 0, 0},
        {1, 1, 0, 1}, 
        {1, 1, 0, 0},
        {0, 1, 1, 1}}

Output: DDRDRR DRDDRR

```py
{!10_recusion_and_backtracking/05_rat_in_maze.py!}
```

[📘](https://takeuforward.org/data-structure/rat-in-a-maze/) [💻](https://www.geeksforgeeks.org/problems/rat-in-a-maze-problem/1)<br>

---

## Word Break (print all ways)

**❓**: You are given a non-empty string S containing no spaces’ and a dictionary of non-empty strings (say the list of words). You are supposed to construct and return all possible sentences after adding spaces in the originally given string ‘S’, such that each word in a sentence exists in the given dictionary.

Note :
The same word in the dictionary can be used multiple times to make sentences.
Assume that the dictionary does not contain duplicate words.

**Example**:  
1  
6  
god is now no where here  
godisnowherenowhere  

Output:  
god is no where no where  
god is no where now here  
god is now here no where  
god is now here now here  
=== "🐍"
    ```py
    --8<-- "10_recusion_and_backtracking/06_word_break.py"
    ```
=== "🦀"
    ```rust
    --8<-- "10_recusion_and_backtracking/06_word_break.rs"
    ```
[💻](https://www.naukri.com/code360/problems/983635?topList=striver-sde-sheet-problems&leftPanelTabValue=PROBLEM)<br>

---