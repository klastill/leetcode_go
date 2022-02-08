package leetcode

func addDigits(num int) int {
	i := 0
	n := 1000
	for i < n {
		sum := 0
		for num != 0 {
			sum += num % 10
			num /= 10
		}
		num = sum

		if num < 10 {
			break
		}
	}

	return num
}

// test case:
// 10001
// 38

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Digits.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Add Digits.
