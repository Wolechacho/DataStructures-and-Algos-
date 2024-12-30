package Arrays

import "fmt"

func MissingNumberProgression(nums []int) int {
	d := (nums[len(nums)-1] - nums[0]) / len(nums)
	fmt.Println(d)
	c := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+d {
			return nums[i-1] + d
		}
	}
	
	return c
}

func MissingNumberProgression1(arr []int)int{
	n := len(arr)
	
	difference := (arr[n-1] - arr[0]) / n
	fmt.Println("difference",difference)
	 i  := arr[0]
	 
	 for _,num := range arr{
		 fmt.Println("num",num)
		if  num != i{
			return i
		}
		i+= difference
		 fmt.Println("i",i)
	 }
	 
	 return i
}
