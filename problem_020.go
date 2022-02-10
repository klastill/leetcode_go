package leetcode

import "strings"

func isValid(s string) bool {
	answer := true
	var stack []string
	tosl := strings.Split(s, "")
	for _, c := range tosl {
		if c == "(" || c == "{" || c == "[" {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				answer = false
				break
			}
			if c == ")" {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				} else {
					answer = false
					break
				}
			} else if c == "}" {
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				} else {
					answer = false
					break
				}
			} else if c == "]" {
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				} else {
					answer = false
					break
				}
			}
		}
	}
	if len(stack) != 0 {
		answer = false
	}
	return answer
}

// test case:
// "()"
// "{{}}}"
// "{{({))}}"

// Runtime: 4 ms, faster than 17.97% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.3 MB, less than 16.38% of Go online submissions for Valid Parentheses.
