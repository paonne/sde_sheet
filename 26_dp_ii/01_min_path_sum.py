class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        return self.helper(grid, 0, 0, {})

    def helper(self, grid, r, c, dp):
        if (r, c) in dp:
            return dp[(r, c)]

        if r == len(grid) - 1 and c == len(grid[0]) - 1:
            return grid[r][c]

        if r == len(grid) or c == len(grid[0]):
            return float("inf")

        p1 = self.helper(grid, r + 1, c, dp)
        p2 = self.helper(grid, r, c + 1, dp)

        dp[(r, c)] = grid[r][c] + min(p1, p2)
        return dp[(r, c)]
