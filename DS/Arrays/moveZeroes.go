package Arrays

import "fmt"

func moveZeroes(nums []int) {
	frontIndex := 0
	rearIndex := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if frontIndex >= rearIndex {
			break
		}
		if nums[i] == 0 {
			for j := frontIndex; j < rearIndex; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			rearIndex--
			fmt.Println(nums)
		}
		frontIndex++
	}
}

func moveZeroes1(nums []int) {

	lastZeroIndex := 0
	 for i := 0; i < len(nums);i++{
        if nums[i] != 0{
            nums[i],nums[lastZeroIndex] = nums[lastZeroIndex],nums[i]
            lastZeroIndex++
        }
        
    }

	fmt.Println(nums)
}
