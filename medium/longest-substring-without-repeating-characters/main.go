package main

// runtime: 127ms beats 17.22%
// memory: 9.39MB beats 6.98%
func lengthOfLongestSubstring(s string) int {
	max := nUnique(s)

	for i := len(s) - max; i < len(s); i++ {
		for j := 0; j < i+1; j++ {
			subs := s[j : len(s)-i+j]

			if hasNoDuplicates(subs) {
				return len(subs)
			}
		}
	}

	return 0
}

func hasNoDuplicates(s string) bool {
	m := make(map[rune]struct{}, len(s))

	for _, c := range s {
		if _, exists := m[c]; !exists {
			m[c] = struct{}{}
		} else {
			return false
		}
	}

	return true
}

func nUnique(s string) int {
	m := make(map[rune]struct{}, len(s))

	for _, c := range s {
		if _, exists := m[c]; !exists {
			m[c] = struct{}{}
		}
	}

	return len(m)
}

func main() {
	_ = lengthOfLongestSubstring("abc")
}
