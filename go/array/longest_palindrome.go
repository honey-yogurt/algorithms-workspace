// 给你一个字符串 s，找到 s 中最长的回文子串 。
package array

// 解决该问题的核心是从中心向两端扩散的双指针技巧。
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var res string
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		// res = longest(res, s1, s2)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

// 如果回文串的长度为奇数，则它有一个中心字符；如果回文串的长度为偶数，则可以认为它有两个中心字符。所以我们可以先实现这样一个函数：
// 在 s 中寻找以 s[l] 和 s[r] 为中心的最长回文串
func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 双指针，向两边展开
		l--
		r++
	}
	// 返回以 s[l] 和 s[r] 为中心的最长回文串
	return s[l+1 : r]
}
