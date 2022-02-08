package leetcode

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var answer float64
	var merged []int
	merged = append(merged, nums1...)
	merged = append(merged, nums2...)
	sort.Slice(merged, func(i, j int) bool {
		return merged[i] < merged[j]
	})
	fmt.Println(merged)
	if len(merged)%2 == 0 {
		answer = float64(merged[len(merged)/2-1]+merged[len(merged)/2]) / float64(2)
	} else {
		answer = float64(merged[len(merged)/2])
	}
	return answer
}

//test case:
//[1,3] [4]
//[1,2,3] [4]

// Runtime: 24 ms, faster than 37.00% of Go online submissions for Median of Two Sorted Arrays.
// Memory Usage: 6.4 MB, less than 28.06% of Go online submissions for Median of Two Sorted Arrays.
