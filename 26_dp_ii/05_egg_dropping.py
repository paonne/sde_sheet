class Solution:
    def eggDrop(self, n, k):
        return self.helper(n, k, {})

    def helper(self, n, k, dp):
        if (n, k) in dp:
            return dp[(n, k)]

        if n == 1 or k <= 0:
            return k

        res = float("inf")
        for i in range(1, k + 1):
            not_break = self.helper(n, k - i, dp)
            breaks = self.helper(n - 1, i - 1, dp)
            res = min(res, 1 + max(not_break, breaks))

        dp[(n, k)] = res
        return dp[(n, k)]
