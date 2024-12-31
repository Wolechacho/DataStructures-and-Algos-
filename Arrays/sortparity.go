package Arrays

import "fmt"

func SortParity(nums []int) {
	front := 0
	rear := len(nums) - 1
	output := make([]int, len(nums))
	i := 0
	for front <= rear {
		if nums[i]%2 == 0 {
			output[front] = nums[i]
			front++
		} else {
			output[rear] = nums[i]
			rear--
		}
		i++
	}

	fmt.Println(output)
}
