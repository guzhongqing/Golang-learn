package main

import "fmt"

func twoSumFor(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}

}

func twoSum(nums []int, target int) []int {

	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		valve, ok := hashmap[target-nums[i]]

		if ok {
			return []int{i, valve}
		} else {
			hashmap[nums[i]] = i
		}

	}

	return []int{}

}

func 两数之和() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSumFor(nums, target)
	fmt.Println(ans)
}
