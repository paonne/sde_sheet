class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        res = float("-inf")

        pre, suff = 1, 1
        n = len(nums)

        for i in range(n):
            pre = pre * nums[i]
            suff = suff * nums[n - i - 1]

            res = max(pre, suff, res)

            if pre == 0:
                pre = 1
            if suff == 0:
                suff = 1

        return res
