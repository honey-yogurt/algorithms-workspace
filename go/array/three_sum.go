// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
package array

import (
	"sort"
)

func threeNum(nums []int) [][]int {
	n := len(nums)
	// 排序，从小到大
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 固定 k ,双指针从 k+1(i) 和 最后一个(j)开始
	// k 直接固定到倒数第三个，留 i j 双指针，如果不满足就可以结束循环了
	for k := 0; k < n-2; k++ {
		// 如果开头的元素 > 0，后续所有的元素都大于0，直接结束循环
		if nums[k] > 0 {
			break
		}
		// 如果 k 同上一个一样，说明 i j 也会一样，直接跳过这个值
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		// 如果前三个（最小的三个）已经大于0了，后续所有k就没必要再循环了（因为k只会越来越大）
		if nums[k]+nums[k+1]+nums[k+2] > 0 {
			break
		}

		// 如果对于当前k ，最大的和还小于0，当前k也没必要循环了
		if nums[k]+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		// 双指针
		i, j := k+1, n-1
		// 双指针结束条件 i < j
		for i < j {
			s := nums[k] + nums[i] + nums[j]
			if s == 0 {
				// 记录下标，i j 同时移动
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				// 去重 k 相同，如果 i j 有一个重复，另一个必定重复，所以 i j 都要去重

				// i == i-1， 执行i++
				for i++; i < j && nums[i] == nums[i-1]; i++ {
					// 循环内什么都不用干，只要++就行
				}

				for j--; i < j && nums[j] == nums[j+1]; j-- {

				}
			} else if s < 0 {
				i++
			} else {
				j--
			}

		}
	}
	return ans
}
