class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        return self.helper(word1, word2, dp={})

    def helper(self, word1, word2, ind1=0, ind2=0, dp={}):
        if (ind1, ind2) in dp:
            return dp[(ind1, ind2)]

        if ind1 == len(word1) and ind2 == len(word2):
            return 0

        if ind1 == len(word1):
            return len(word2) - ind2

        if ind2 == len(word2):
            return len(word1) - ind1

        if word1[ind1] == word2[ind2]:
            dp[(ind1, ind2)] = self.helper(word1, word2, ind1 + 1, ind2 + 1, dp)
            return dp[(ind1, ind2)]

        insert = self.helper(word1, word2, ind1, ind2 + 1, dp)
        delete = self.helper(word1, word2, ind1 + 1, ind2, dp)
        replace = self.helper(word1, word2, ind1 + 1, ind2 + 1, dp)

        dp[(ind1, ind2)] = 1 + min(insert, delete, replace)
        return dp[(ind1, ind2)]
