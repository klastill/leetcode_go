package leetcode

func myPow(x float64, n int) float64 {
	powed := 1.
	if x == 1 {
		return x
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}
	if n%2 == 0 {
		n = n / 2
		x *= x
	}
	if n > -1 {
		for i := 0; i < n; i++ {
			powed *= x
		}
	} else {
		for i := 0; i > n; i-- {
			powed *= x
		}
		powed = 1 / powed
	}

	return powed
}

// test case:
// -1. -2100000
//	2.3 7

// Runtime: 3896 ms, faster than 5.01% of Go online submissions for Pow(x, n).
// Memory Usage: 2.1 MB, less than 76.49% of Go online submissions for Pow(x, n).
