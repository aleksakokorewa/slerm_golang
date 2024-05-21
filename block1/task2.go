package main

import (
	"fmt"
	"os"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 3, 5, 4, 3}
	nums3 := []int{}
	fmt.Println(dubs(nums))
	fmt.Println(dubs(nums2))
	fmt.Println(dubs(nums3))
}

func dubs(nums []int) bool {
	if len(nums) <= 1 {
		fmt.Println("Empty array")
		os.Exit(1)
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
