package Arrays

import "fmt"

func CheckMajorityElement(nums []int, target int) bool {

	// if len(nums) == 1 {
	// 	return nums[0] == target
	// }

	// if len(nums) == 2 {
	// 	return nums[0] == target && nums[1] == target
	// }

	// med := nums[len(nums)/2-1]
	// med1 := nums[len(nums)/2+1]
	// return med == target && med1 == target

	i := 0

	for i < len(nums) && nums[i] != target {
		fmt.Println(nums[i])
		i++
	}
	j := i

	for j < len(nums) && nums[j] == target {
		j++
	}
	fmt.Println("i", i)
	fmt.Println("j", j)
	return (j - i) > len(nums)/2
}

// func hdhd(){
// 	   m := map[int]int{}

//     for _,num := range nums{
//         m[num]++
//     }

//     max := 0
//     val := 0

//     for k,v := range m{
//         if v > max{
//             max = v
//             val = k
//         }
//     }

//     if max > len(nums) / 2 && val == target{
//         return val
//     }
//     return -1
// }
