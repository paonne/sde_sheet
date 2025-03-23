class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        return self.helper(s, set(wordDict), 0, {})

    def helper(self, s, words, ind, dp):
        if ind in dp:
            return dp[ind]

        if ind == len(s):
            return True

        for i in range(ind + 1, len(s) + 1):
            if s[ind:i] in words and self.helper(s, words, i, dp):
                dp[ind] = True
                return dp[ind]

        dp[ind] = False
        return dp[ind]
