package binarysearch

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursive(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	return binarySearch(arr, 0, len(arr), target)
}

func binarySearch(arr []int, low, high, target int) int {
	// 查找到最后一个元素，还是没有找到 target ，此时 low = high = mid ,
	// 然后继续迭代，然后要么就是[mid+1,high]，要么就是[low,high-1]，
	// 此时应该退出迭代了，因为数组中已经没有满足条件的值了。
	// 故退出条件是 low > high
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearch(arr, mid+1, high, target)
	} else {
		return binarySearch(arr, low, mid-1, target)
	}
}

func BinarySearchFirst(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			// 如果mid等于0，那这个元素已经是数组的第一个元素，那它肯定是我们要找的；如果mid不等于0，但a[mid]的前一个元素a[mid-1]不等于value，那也说明a[mid]就是我们要找的第一个值等于给定值的元素。
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				// 如果经过检查之后发现a[mid]前面的一个元素a[mid-1]也等于value，那说明此时的a[mid]肯定不是我们要查找的第一个值等于给定值的元素。那我们就更新high=mid-1，因为要找的元素肯定出现在[low, mid-1]之间。
				high = mid - 1
			}
		}
	}
	return -1
}

func BinarySearchLast(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			// 最后一个或者后一个
			if mid == len(arr)-1 || arr[mid+1] != target {
				return mid
			} else {
				// 肯定是在后半区
				low = mid + 1
			}
		}
	}
	return -1
}

func BinarySearchFirstMore(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] >= target {
			if mid == 0 || arr[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchLastLess(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] <= target {
			if mid == len(arr)-1 || arr[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
