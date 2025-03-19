class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        res = self.helper(coins, amount, 0, {})
        return res if res != float("inf") else -1

    def helper(self, coins, amount, ind, dp):
        if (amount, ind) in dp:
            return dp[(amount, ind)]

        if amount == 0:
            return 0

        if ind == len(coins):
            return float("inf")

        p1 = self.helper(coins, amount, ind + 1, dp)
        p2 = float("inf")
        if coins[ind] <= amount:
            p2 = 1 + self.helper(coins, amount - coins[ind], ind, dp)

        dp[(amount, ind)] = min(p1, p2)
        return dp[(amount, ind)]
