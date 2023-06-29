package main

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else if target > nums[len(nums)-1] {
			return len(nums)
		} else if target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}
	return 0
}

// func main() {
// 	nums := []int{1, 3, 5, 6}
// 	target := 5
// 	searchInsert(nums, target)
// }
