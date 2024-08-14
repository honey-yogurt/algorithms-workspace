// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
package array

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	// slow 代表下一个不为0的位置
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// [:slow-1]都是非0元素
	for i := slow; i < fast; i++ {
		nums[i] = 0
	}
}
