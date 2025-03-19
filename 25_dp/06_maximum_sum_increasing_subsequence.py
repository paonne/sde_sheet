class Solution:
    def maxSumIS(self, arr):
        n = len(arr)
        dp = arr[:]

        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                if arr[i] < arr[j]:
                    dp[i] = max(dp[i], arr[i] + dp[j])

        return max(dp)
