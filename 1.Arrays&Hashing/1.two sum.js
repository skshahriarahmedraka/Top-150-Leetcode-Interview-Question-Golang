var twoSum = function(nums, target) {
    let m={}

    for (i = 0; i < nums.length; i++){
        // j = nums[i]
        // console.log("\U0001f680 ~ file: two sum.js ~ line 14 ~ twoSum ~ j", j)
        // console.log("\U0001f680 ~ file: two sum.js ~ line 16 ~ twoSum ~ m[target - j]", m[target - j])
        if (m[target - nums[i]]!== undefined) {
            x=m[target-nums[i]]
            // console.log("\U0001f680 ~ file: two sum.js ~ line 17 ~ twoSum ~ x", x)
            return [x,i]
        }
        m[nums[i]]=i

    }
};

// z=[2,7,11,15]
// t=9

// twoSum(z,t)

// console.log("\U0001f680 ~ file: two sum.js ~ line 31 ~ twoSum(z,t)", twoSum(z,t))

