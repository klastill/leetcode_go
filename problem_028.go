package leetcode

func strStr(haystack string, needle string) int {
	answer := -1

	if needle == "" {
		return 0
	}

	i := 0
	for i <= len(haystack)-len(needle) {
		part := haystack[i : i+len(needle)]
		if part == needle {
			answer = i
			break
		}
		i++
	}

	return answer
}

// test case:
// "hello" "ll"

// Runtime: 4 ms, faster than 87.37% of Go online submissions for Implement strStr().
// Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Implement strStr().
