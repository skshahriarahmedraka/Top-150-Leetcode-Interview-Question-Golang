var twoSum = function(nums, target) {
   
    let obj = {
    }

    for(let i = 0; i < nums.length; i++) {
        if(obj[target - nums[i]] !== undefined) {
            return([i, obj[target - nums[i]]] + 1) 
        }
        obj[nums[i]] = i
    }
    
};


