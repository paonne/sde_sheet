class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        res = []
        n = len(nums)
        nums.sort()
        for i in range(n):
            if i != 0 and nums[i] == nums[i - 1]:
                continue
            for j in range(i + 1, n):
                if j != i + 1 and nums[j] == nums[j - 1]:
                    continue
                left, right = j + 1, n - 1
                required_sum = target - nums[i] - nums[j]
                while left < right:
                    if nums[left] + nums[right] < required_sum:
                        left += 1
                    elif nums[left] + nums[right] > required_sum:
                        right -= 1
                    else:
                        res.append([nums[i], nums[j], nums[left], nums[right]])
                        left += 1
                        right -= 1
                        while left < right and nums[left] == nums[left - 1]:
                            left += 1
                        while left < right and nums[right] == nums[right + 1]:
                            right -= 1

        return res


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        res = set()
        n = len(nums)
        nums.sort()
        for i in range(n):
            for j in range(i + 1, n):
                left, right = j + 1, n - 1
                required_sum = target - nums[i] - nums[j]
                while left < right:
                    if nums[left] + nums[right] < required_sum:
                        left += 1
                    elif nums[left] + nums[right] > required_sum:
                        right -= 1
                    else:
                        res.add((nums[i], nums[j], nums[left], nums[right]))
                        left += 1
                        right -= 1
        return list(res)
