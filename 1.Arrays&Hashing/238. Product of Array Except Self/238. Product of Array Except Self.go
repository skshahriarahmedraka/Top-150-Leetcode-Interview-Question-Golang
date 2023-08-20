package main

func main() {

}

func productExceptSelf(nums []int) []int {
	ans:= make([]int,len(nums))
	for i:=1 ;i<=len(nums)-1;i++ {
		if i== 1 {
			ans[i]=nums[i-1]
		}else {
			ans[i]=ans[i-1]*nums[i-1]
		}
	}
	
	left:= nums[len(nums)-1]
	for i:=len(nums)-2;i>=0;i--{
		if (i==0) {ans[i]=left
		} else{
			ans[i]*=left
			left*=nums[i]
		}
	}
	return ans  
    
}