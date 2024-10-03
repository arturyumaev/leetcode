package main

import "strconv"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	result := "1"
	for i := 0; i < n-1; i++ {
		result = rle(result)
	}

	return result
}

// runtime: 3ms beats 70.56%
// memory: 13.19MB beats 6.36%
func rle(s string) string {
	curr := s[0]
	c := 0

	result := ""

	for i := 0; i < len(s); i++ {
		if curr == s[i] {
			c++
			continue
		}

		result += strconv.Itoa(c)
		result += string(curr)

		c = 1
		curr = s[i]
	}
	result += strconv.Itoa(c)
	result += string(curr)

	return result
}

func main() {
	countAndSay(0)
}
