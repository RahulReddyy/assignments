package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	map1 := map[int]string{}
	for _, value := range nums {
		map1[value] = ""
	}
	nums = nums[0:0]
	for key := range map1 {
		nums = append(nums, key)
	}
	sort.Ints(nums)
	return len(nums)
}

func main4() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))

}
