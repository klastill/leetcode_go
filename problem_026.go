package leetcode

func removeDuplicates(nums []int) int {
	i := 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

// test case:
// [1,1,2]
// [0,0,1,1,1,2,2,3,3,4]

// Runtime: 72 ms, faster than 7.28% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.5 MB, less than 56.84% of Go online submissions for Remove Duplicates from Sorted Array.
