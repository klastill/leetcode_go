package leetcode

func twoSum(nums []int, target int) []int {
	var answer []int
	for i, first := range nums {
		for j, second := range nums[i+1:] {
			if first+second == target {
				answer = append(answer, i, i+j+1)
			}
		}
	}
	return answer
}

//test case:
//[2,7,11,15] 9
// [3,2,4] 6

// Runtime: 44 ms, faster than 19.30% of Go online submissions for Two Sum.
// Memory Usage: 3.6 MB, less than 94.70% of Go online submissions for Two Sum.
