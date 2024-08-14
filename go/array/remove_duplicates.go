// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
//
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
//
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 返回 k 。
package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针，快指针用来遍历原数组，慢指针用来维护处理后的新数组
	slow, fast := 0, 0
	// 快指针走到头就结束了
	for fast < len(nums) {
		// 事先已经排序了，非严格递增排列
		// 或者这里用 !=
		if nums[fast] > nums[slow] {
			// 慢指针是处理过后的值
			slow++
			// 修改原数组
			// [:slow] 是我们处理后的数组
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 这个优化的最终在leetcode中执行时间反而更加长了
// 可能是leetcode的测试数组没有这种不重复的数组，导致根本就没有节省复制的成本，反而增加了每次循环的判断成本（fast-slow > 1）
// 当然，如果数组中数据结构是一个复制成本很大的结构，这种优化肯定是有效的
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针，快指针用来遍历原数组，慢指针用来维护处理后的新数组
	slow, fast := 0, 0
	// 快指针走到头就结束了
	for fast < len(nums) {
		// 事先已经排序了，非严格递增排列
		// 或者这里用 !=
		if nums[fast] > nums[slow] {
			if fast-slow > 1 {
				// 慢指针是处理过后的值
				slow++
				// 修改原数组
				// [:slow] 是我们处理后的数组
				nums[slow] = nums[fast]
			} else {
				// 如果每个数组没有重复的数组，不增加 fast - slow == 1 判断，每次都要重复复制一下 fast
				slow++
			}

		}
		fast++
	}
	return slow + 1
}
