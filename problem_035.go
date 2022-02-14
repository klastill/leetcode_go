package leetcode

func searchInsert(nums []int, target int) int {
	answer := len(nums)
	for i, val := range nums {
		if val >= target {
			answer = i
			break
		}
	}
	return answer
}

// test case:
// "[1,3,5,6] 5
// [1,3,5,6] 7

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search Insert Position.
// Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Search Insert Position.
