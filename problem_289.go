package leetcode

func findTheDifference(s string, t string) byte {
	bs := []byte(s)
	bt := []byte(t)

	i := 0
	n := len(bs)

	for i < n {
		val := bs[i]
		if indexOf(val, bt) > -1 {
			bs = append(bs[:indexOf(val, bs)], bs[indexOf(val, bs)+1:]...)
			bt = append(bt[:indexOf(val, bt)], bt[indexOf(val, bt)+1:]...)
			n--
			i--
		}
		i++
	}
	return bt[0]
}

func indexOf(s byte, slice []byte) int {
	for i, val := range slice {
		if s == val {
			return i
		}
	}
	return -1
}

// test case:
// "a" "aa"
// "abcd" "abcde"
// "aabbc" "aabbbc"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Find the Difference.
// Memory Usage: 2.1 MB, less than 95.41% of Go online submissions for Find the Difference.
