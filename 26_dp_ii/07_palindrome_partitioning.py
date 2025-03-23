class Solution:
    def palindromicPartition(self, s):
        return self.helper(s, 0, {})

    def helper(self, s, ind, dp):
        if ind in dp:
            return dp[ind]

        if s[ind:] == s[ind:][::-1]:
            return 0

        res = len(s)

        for i in range(ind + 1, len(s)):
            curr_string = s[ind:i]

            if curr_string == curr_string[::-1]:
                res = min(res, 1 + self.helper(s, i, dp))

        dp[ind] = res
        return dp[ind]
