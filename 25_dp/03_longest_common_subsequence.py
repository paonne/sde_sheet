class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        return self.helper(text1, text2, dp={})

    def helper(self, text1, text2, ind1=0, ind2=0, dp={}):
        if (ind1, ind2) in dp:
            return dp[(ind1, ind2)]

        if ind1 == len(text1) or ind2 == len(text2):
            return 0

        if text1[ind1] == text2[ind2]:
            dp[(ind1, ind2)] = 1 + self.helper(text1, text2, ind1 + 1, ind2 + 1, dp)
            return dp[(ind1, ind2)]

        p1 = self.helper(text1, text2, ind1, ind2 + 1, dp)
        p2 = self.helper(text1, text2, ind1 + 1, ind2, dp)

        dp[(ind1, ind2)] = max(p1, p2)
        return dp[(ind1, ind2)]
