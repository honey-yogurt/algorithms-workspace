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

func BucketSort(arr []int) {
	num := len(arr)
	if num < 1 {
		return
	}
	max := Max(arr)
	// 二维切片
	buckets := make([][]int, num) // 最大为 num，万一一个桶一个元素呢，每个元素就是桶
	index := 0                    // 桶序号
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max // arr[i] / max = [0~1]，然后乘以 num-1，按比例尽可能均匀分配到桶上，保证桶之间是有序的
		buckets[index] = append(buckets[index], arr[i])
	}
	pos := 0 // 标记数组的位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])       // 依次快排
			copy(arr[pos:], buckets[i]) // 排序写回原数组
			pos += bucketLen
		}
	}
}

func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	max := Max(arr)

	c := make([]int, max+1) // 从 0 到 max+1，下标索引就代表这个 要排序的元素值
	// 这个循环结束，c 每个下标对应的值 就代表 arr中元素值等于下标的元素有多少个
	// 比如c[0] 代表在 待排序的数组（arr）中，0 有多少个，a[max] 代表 最大值有多少个
	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}
	// 此时我们只知道 arr 中给一个元素，该元素有多少个，但是并不知道这个元素排序位置，也就是在这个元素之前有多少个

	// 对c 进行顺序累加，这时候给定一个元素m，c[m]就代表 小于等于 c[m] 有多少个了。
	// 但是仍然不知道该元素的排名，因为此时已经不知道该元素有多少个了，并且这个元素在一堆相同的元素中的排序
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	// 倒序扫描 arr，每扫到一个元素m，该m肯定是所以一堆m 中最大的，也就是它的排序就是 c[m]，然后 c[m]-1，下一次扫到 m 时候，这个 m 就是c[m]
	r := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := c[arr[i]] - 1 // arr[i] 排序是 c[i]，下标 -1
		r[index] = arr[i]      // 放入对应的序号
		c[arr[i]]--            // 累加减一
	}
	copy(arr, r)
}

func RadixSort(arr []int) {
	// 最大值，判断总共多少位
	max := Max(arr)
	// 从最低位（个位）开始处理
	exp := 1
	// 向下取整，第一次循环对个位排序，第二次循环对十位排序 ...
	for int(max/exp) > 0 {
		arr = countingSortByExp(arr, exp)
		exp *= 10
	}
}

func countingSortByExp(arr []int, exp int) []int {
	n := len(arr)
	// 输出排序好的临时结果
	output := make([]int, n)

	// 0-9 共 10 位
	count := make([]int, 10)
	// 第一遍循环得到每个数字的 count，注意位数
	for i := 0; i < n; i++ {
		expNum := (arr[i] / exp) % 10 // 所在 exp 位的数字
		count[expNum]++
	}
	// 顺序累加
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// 倒序遍历 arr
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	copy(arr, output)
	return arr
}

func Max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
