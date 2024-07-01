package sort

func BubbleMaxSort1(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	for i := 0; i < n; i++ {
		// 每一轮都少比较一个。因为每一轮都确定了一个最大位置
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tem := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tem
			}
		}
	}
	return arr
}

func BubbleMaxSort2(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	for i := 0; i < n; i++ {
		flag := false // 提前退出冒泡排序的标志，默认没有发生交换
		// 每一轮都少比较一个。因为每一轮都确定了一个最大位置
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tem := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tem
				// 发生交换
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// 怎么写成了 O(n^3)
func InsertionSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			// 保证稳定性，等于也要后移
			if arr[i] < arr[j] {
				// 插入j此刻的位置，j及以后的都要移动以为，让 i 插入
				// 存储 i，这样移动时候可以直接覆写i
				tem := arr[i]
				for k := i; k > j; k-- {
					arr[k] = arr[k-1]
				}
				arr[j] = tem
				break
			}
		}
	}
	return arr
}

func InsertionSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		// 待插入的值
		val := arr[i]
		// 待比较的终点
		j := i - 1
		// 从右往左一遍遍历，一边挪位置
		for ; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = val
	}
	return arr
}

func SelectSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// 第i轮交换，确定arr[i]最小
	for i := 0; i < len(arr); i++ {
		m := arr[i]
		pos := i
		for j := i; j < len(arr); j++ {
			if arr[j] < m {
				m = arr[j]
				pos = j
			}
		}
		arr[pos] = arr[i]
		arr[i] = m
	}

	return arr
}
