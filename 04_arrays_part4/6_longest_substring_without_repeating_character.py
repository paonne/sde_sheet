class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        n = len(s)
        if n == 0:
            return 0
        left = res = 0
        seen = {}
        for i in range(n):
            if s[i] in seen and seen[s[i]] >= left:
                res = max(res, i - left)
                left = seen[s[i]] + 1
            seen[s[i]] = i

        return max(res, i - left + 1)
