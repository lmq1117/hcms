// Package word provides utilities for word games.
package word

// IsPalindrome reports whether s reads the same forward and backward.
// 检查一个字符串是否从前向后和从后向前读都是一样的
// (Our first attempt.)
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
