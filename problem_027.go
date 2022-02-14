package leetcode

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}

// test case:
// [3,2,2,3] 3

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
// Memory Usage: 2.1 MB, less than 82.10% of Go online submissions for Remove Element.
