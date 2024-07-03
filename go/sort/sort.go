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
				//tem := arr[j]
				//arr[j] = arr[j+1]
				//arr[j+1] = tem
				arr[j], arr[j+1] = arr[j+1], arr[j]
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

// 这种直接对原始数组操作，可能不是很好理解
func MergeSort1(arr []int) {
	p := 0
	q := len(arr) - 1
	mergeSort1(arr, p, q)
}

func mergeSort1(arr []int, l, r int) {
	if l == r {
		return
	}
	// 默认向下取整
	m := (l + r) / 2
	mergeSort1(arr, l, m)
	mergeSort1(arr, m+1, r)
	merge1(arr, l, m, r)
}

func merge1(arr []int, l, m, r int) {
	temp := make([]int, r-l+1)
	// 双指针
	i := l
	j := m + 1

	// 临时数组下标
	k := 0
	for ; i <= m && j <= r; k++ {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
	}
	// 其中一组（不知道哪一组）遍历结束

	// 左边没有遍历完
	for ; i <= m; i++ {
		temp[k] = arr[i]
		k++
	}

	// 右边没有遍历完
	for ; j <= r; j++ {
		temp[k] = arr[j]
		k++
	}
	copy(arr[l:r+1], temp)
}

// 返回的不是原数组了，但是易于理解
func MergeSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	m := len(arr) / 2
	left := MergeSort2(arr[:m])
	right := MergeSort2(arr[m:])
	// 排序
	return merge2(left, right)
}

func merge2(left, right []int) []int {
	tem := make([]int, len(left)+len(right))
	i, j := 0, 0
	k := 0
	m, n := len(left), len(right)
	for {
		if i >= m || j >= n {
			break
		}
		if left[i] <= right[j] {
			tem[k] = left[i]
			i++
		} else {
			tem[k] = right[j]
			j++
		}
		k++
	}
	for ; i < m; i++ {
		tem[k] = left[i]
		k++
	}
	for ; j < n; j++ {
		tem[k] = right[j]
		k++
	}
	return tem
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	q := partition(arr, l, r)
	quickSort(arr, l, q-1)
	quickSort(arr, q+1, r)
}

func partition(arr []int, l, r int) int {
	pivot := arr[r]
	i := l // 最终 pivot 所在的索引,，运行时指代的是未分区的第一索引位置
	// 每次把 [j,r-1] 中默认都是没有比较的值，[l,i-1]都是处理比较好的数据（也就是小于pivot）
	for j := l; j < r; j++ {
		// 需要交换到 arr[j]，这样 arr[j] 就小于 pivot，未处理区间缩小，i++
		if arr[j] < pivot {
			if i != j {
				//tem := arr[j]
				//arr[j] = arr[i]
				//arr[i] = tem
				arr[i], arr[j] = arr[j], arr[i]
			}
			// 收缩未处理分区
			i++
		}
	}
	// 最后把 r（pivot）交换到 arr[i]，返回 i
	//tem := arr[i]
	//arr[i] = pivot
	//arr[r] = tem
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
