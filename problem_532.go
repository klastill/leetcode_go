package leetcode

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	answer := 0
	var set []int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)

	i := 0

	for i < len(nums)-1 {
		j := i + 1
		for j < len(nums) {
			if (nums[j]-nums[i]) == k && !isInSlice((nums[j]+nums[i]), set) {
				set = append(set, (nums[j] + nums[i]))
				answer++
			}
			j++
		}
		i++
	}

	return answer
}

func isInSlice(num int, slice []int) bool {
	for _, i := range slice {
		if num == i {
			return true
		}
	}
	return false
}

// test case:
// [3,1,4,1,5] 2

// Runtime: 236 ms, faster than 15.56% of Go online submissions for K-diff Pairs in an Array.
// Memory Usage: 5.7 MB, less than 60.00% of Go online submissions for K-diff Pairs in an Array.
