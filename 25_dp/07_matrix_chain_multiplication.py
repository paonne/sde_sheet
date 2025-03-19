class Solution:
    def matrixMultiplication(self, nums):
        return self.helper(nums, 1, len(nums) - 1, {})

    def helper(self, nums, i, j, dp):
        if (i, j) in dp:
            return dp[(i, j)]

        if i == j:
            return 0

        mini = float("inf")

        for k in range(i, j):
            mini = min(
                mini,
                self.helper(nums, i, k, dp)
                + self.helper(nums, k + 1, j, dp)
                + (nums[i - 1] * nums[k] * nums[j]),
            )

        dp[(i, j)] = mini
        return dp[(i, j)]
