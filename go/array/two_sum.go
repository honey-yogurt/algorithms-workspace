package array

// O(n^2) 时间复杂度
// 25ms 3.3MB
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// O(n) 时间复杂度
// 6ms 5.2MB
func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int)
	// 将第二层循环中寻找 target-x 的 O(n) 变成 O(1)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		pos, ok := numsMap[target-nums[i]]
		// 避免将 target/2 的索引输出两遍
		if ok && pos != i {
			return []int{pos, i}
		}
	}
	return nil
}

// O(n) 时间复杂度，但是比 twoSum2 更快
// 5ms 4MB
func twoSum3(nums []int, target int) []int {
	numsMap := map[int]int{}
	// 一遍循环，完成map和校验
	// twoSum2 先是全部存储，势必会出现 nums[i] = target/2 ，num[i] 和 map[target/2]=i ，需要去重
	// twoSum3 并不需要全部存储，map中找不到匹配的就将自己存进map，等待后面的元素来匹配
	for i, num := range nums {
		want := target - num
		pos, ok := numsMap[want]
		if ok {
			return []int{pos, i}
		}
		numsMap[num] = i
	}
	return nil
}

// 3ms 4MB
func twoSum4(nums []int, target int) []int {
	numsMap := map[int]int{}
	for i, num := range nums {
		// 少了一次变量赋值
		// 直接在 if 语句中进行判断和赋值，可能导致更好的分支预测和编译器优化？
		if pos, ok := numsMap[target-num]; ok {
			return []int{pos, i}
		}
		numsMap[num] = i
	}
	return nil
}
