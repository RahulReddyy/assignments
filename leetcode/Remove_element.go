package main

import "fmt"

func removeElement(nums []int, val int) int {
	nums1 := []int{}
	nums1 = append(nums1, nums...)
	nums = nums[0:0]

	for _, value := range nums1 {
		if value == val {
			continue
		}
		nums = append(nums, value)
	}
	fmt.Println(nums)
	return len(nums)
}
func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	removeElement(nums, val)
}
