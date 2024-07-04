package binarysearch

import (
	"fmt"
	"testing"
)

// 有序
var pending = []int{14, 23, 89, 123, 256, 374, 650, 999}
var pendingRepeated = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}

func TestBinarySearch(t *testing.T) {
	fmt.Println(BinarySearch(pending, 123))                // 3
	fmt.Println(BinarySearchRecursive(pending, 123))       // 3
	fmt.Println(BinarySearchFirst(pendingRepeated, 8))     // 5
	fmt.Println(BinarySearchLast(pendingRepeated, 8))      // 7
	fmt.Println(BinarySearchFirstMore(pendingRepeated, 7)) // 5
	fmt.Println(BinarySearchLastLess(pendingRepeated, 9))  // 7
}
