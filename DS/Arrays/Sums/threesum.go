package Sums

import "sort"
func ThreeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    for i:=0;i < len(nums)-2;i++{
        if ((i == 0) || (i > 0 && nums[i] != nums[i-1])){
            low := i + 1
            high := len(nums) - 1
            sum := 0-nums[i]
            for low < high{
                if nums[low] + nums[high] == sum{
                    arr := []int{}
                    arr = append(arr,nums[i],nums[low],nums[high])
                    result = append(result,arr)
                    for low < high && nums[low] == nums[low+1]{
                        low++
                    }
                    for low < high && nums[high] == nums[high-1]{
                        high--
                    }
                    low++
                    high--
                }else if(nums[low] + nums[high] > sum){
                    high--
                }else{
                    low++
                }
            }
        }
    }
    
    return result
}