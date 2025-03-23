class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        target = sum(nums)
        if target % 2:
            return False

        return self.helper(nums, target // 2, 0, {})

    def helper(self, nums, target, ind, dp):
        if (target, ind) in dp:
            return dp[(target, ind)]

        if target == 0:
            return True

        if target < 0 or ind == len(nums):
            return False

        p1 = self.helper(nums, target, ind + 1, dp)
        p2 = False
        if nums[ind] <= target:
            p2 = self.helper(nums, target - nums[ind], ind + 1, dp)

        dp[(target, ind)] = p1 or p2
        return dp[(target, ind)]
