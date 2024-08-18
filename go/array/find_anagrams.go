// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
package array

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		need[p[i]]++
		// 为什么 need[v] 不行?
	}

	left, right := 0, 0
	valid := 0

	var res []int

	for right < len(s) {
		// 增加窗口
		c := s[right]
		right++

		// 更新窗口值
		// 先找到满足条件的冗余子串
		if _, ok := need[c]; ok {
			window[c]++
			// 其实这时查到的是超集，因为后续会出现 windows[c] > need[c] 的情况，但是我们的子串数量应该是一样大，所以第一步获取的是超集
			if window[c] == need[c] {
				valid++
			}
		}

		// 如果窗口大于 p 的长度了，考虑收缩窗口
		if right-left >= len(p) {
			// 此时说明正好是子串，因为长度一样，满足要求的字符数也一样，必然一样
			if valid == len(need) {
				// 加入起始索引
				res = append(res, left)
			}

			// 此时收缩窗口
			d := s[left]
			left++

			// 更新窗口
			if _, ok := need[d]; ok {
				// 先判断d是否正好 个数相等，如果相等，就要valid--，因为要把这个d移除出去了
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}

	}
	return res
}
