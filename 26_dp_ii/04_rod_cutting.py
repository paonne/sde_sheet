class Solution:
    def minCost(self, n: int, cuts: List[int]) -> int:
        return self.helper(0, n, cuts, {})

    def helper(self, l, r, cuts, dp):
        if (l, r) in dp:
            return dp[(l, r)]

        res = float("inf")
        for cut in cuts:
            if l < cut < r:
                res = min(
                    res,
                    r
                    - l
                    + self.helper(l, cut, cuts, dp)
                    + self.helper(cut, r, cuts, dp),
                )

        dp[(l, r)] = res if res != float("inf") else 0
        return dp[(l, r)]
