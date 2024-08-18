// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
package array

func lengthOfLongestSubstring(s string) int {
	// 定义一个窗口
	// 在窗口中，统计每个字符的数量，存在>1的情况，这个窗口就要收窄
	windows := make(map[rune]int)

	left, right := 0, 0
	// 记录结果
	res := 0

	// 循环条件 窗口右指针 小于 数组长度
	for right < len(s) {
		// 右窗口开始滑动
		c := rune(s[right])
		right++

		// 进行窗口内数据更新
		windows[c] = windows[c] + 1 // num+1

		// 判断左侧窗口是否需要收缩
		// 每次右窗口滑进来一个元素，都要判断该元素数量是否存在>1的情况，存在就说明和前面某个元素重复，
		// 这个窗口就要收窄，因为不知道具体是哪个元素，所以要从左侧挨个收窄
		for windows[c] > 1 {
			// 左侧窗口收窄
			d := rune(s[left])
			left++

			// 更新窗口数据
			windows[d] = windows[d] - 1
		}

		// 在这里更新答案
		// 每次右侧窗口移动，都要比较下窗口大小，记录最大的一次
		if res < (right - left) {
			res = right - left
		}
	}
	return res
}
