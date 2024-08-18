// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。
package array

func checkInclusion(s1 string, s2 string) bool {
	// 首先需要记录 s1 的各个字符的次数
	need := make(map[rune]int)
	for _, s := range s1 {
		need[s]++
	}

	// 创建一个窗口
	window := make(map[rune]int)

	left, right := 0, 0

	// 统计符合的字符数
	valid := 0
	for right < len(s2) {
		// 右移窗口
		c := rune(s2[right])
		right++

		// 更新窗口数据
		// 判断这个元素在不在s1，如果s1中都没有这个元素，此时这个窗口肯定不是正确的
		// 所以就不需要更新窗口了
		if _, ok := need[c]; ok {
			window[c]++
			// 如果当前字符的数量符合其在s1字符的数量，则窗口的有效字符+1
			if window[c] == need[c] {
				valid++
			}
		}

		// 是否需要收缩
		// 至少需要判断窗口大小是否大于等于 s1
		if right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := rune(s2[left])
			left++

			// 更新窗口数据
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 未找到符合条件的子串
	return false
}
