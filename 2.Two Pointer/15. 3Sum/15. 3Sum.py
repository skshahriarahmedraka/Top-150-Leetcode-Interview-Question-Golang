def threeSum(nums):
    if len(nums) < 3:
        return []
    nums.sort()
    res = set()
    for i in range(len(nums) - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        d = {}
        for j in range(i + 1, len(nums)):
            if nums[j] not in d:
                d[-nums[i] - nums[j]] = j
            else:
                res.add((nums[i], -nums[i] - nums[j], nums[j]))
    return list(res)
