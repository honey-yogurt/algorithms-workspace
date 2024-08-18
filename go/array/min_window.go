// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
// 注意：
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
// 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
package array

import (
	"math"
)

func minWindow(s string, t string) string {
	// 1. 统计 t 中每个字符的多少
	// 2. 记录 满足 条件1 的情况的字符串串，然后选择收缩窗口，如果窗口满足条件1的话，继续收缩，然后更新子字符串，
	// 3. 当收缩不满足时候，继续增大窗口，直至满足条件2然后继续条件3，直至窗口到最后

	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	// 创建一个窗口
	windows := make(map[byte]int)

	left, right := 0, math.MaxInt

	i, j := 0, 0
	for j < len(s) {
		// 右移窗口
		c := s[j]
		j++

		// 更新窗口数据
		// 因为要查找最小子串，所以所有数据都要更新，不管这个 c 在不在 t 中
		windows[c]++

		// 判断是否有符合条件的子串出现
		// 缩小窗口
		if valid(need, windows) {

			for i < j {
				d := s[i]
				i++
				// 更新窗口数据
				windows[d]--
				if !valid(need, windows) {
					// 一旦不满足了，就需要重新扩大窗口
					if j-i < right-left {
						left, right = i-1, j
					}
					break
				}

			}
		}

	}
	if right == math.MaxInt {
		return ""
	}
	return s[left:right]
}

func valid(need, window map[byte]int) bool {
	for k, v := range need {
		if window[k] < v {
			return false
		}
	}
	return true
}

// 根本不用每次移动都要判断循环比较是否valid
// 增加一个标识位，每当窗口某个字符达标了，valid 就+1，通过valid 和 t 的长度比较就可以知道是否 符合了
func minWindow1(s string, t string) string {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 扩大窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {

			// 在这里更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// d 是将移出窗口的字符
			d := s[left]
			// 缩小窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}

	}
	// 返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
