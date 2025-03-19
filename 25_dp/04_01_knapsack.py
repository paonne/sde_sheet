class Solution:
    def knapsack01(self, wt, val, n, W):
        return self.helper(wt, val, 0, W, {})

    def helper(self, wt, val, ind, curr, dp):
        if (ind, curr) in dp:
            return dp[(ind, curr)]

        if ind == len(wt):
            return 0

        pick = 0
        if wt[ind] <= curr:
            pick = val[ind] + self.helper(wt, val, ind + 1, curr - wt[ind], dp)

        not_pick = self.helper(wt, val, ind + 1, curr, dp)

        dp[(ind, curr)] = max(pick, not_pick)
        return dp[(ind, curr)]
