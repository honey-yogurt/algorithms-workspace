package sort

import (
	"fmt"
	"testing"
)

func TestBubbleMaxSort(t *testing.T) {
	s := []int{3, 5, 4, 1, 2, 6}
	fmt.Println(BubbleMaxSort1(s))
	fmt.Println(BubbleMaxSort2(s))
	s1 := []int{4, 5, 6, 3, 2, 1}
	fmt.Println(BubbleMaxSort1(s1))
	fmt.Println(BubbleMaxSort2(s1))
	s2 := []int{4}
	fmt.Println(BubbleMaxSort1(s2))
	fmt.Println(BubbleMaxSort2(s2))
}

func TestInsertionSort(t *testing.T) {
	s := []int{3, 5, 4, 1, 2, 6}
	fmt.Println(InsertionSort2(s))
	s = []int{2, 1}
	fmt.Println(InsertionSort2(s))
	s = []int{1}
	fmt.Println(InsertionSort1(s))
	s = []int{}
	fmt.Println(InsertionSort1(s))
}

func TestSelectSort(t *testing.T) {
	s := []int{3, 5, 4, 1, 2, 6, 1, 2, 8, 5, 10}
	fmt.Println(SelectSort(s))
}

func TestMergeSort(t *testing.T) {
	s := []int{3, 5, 4, 1, 2, 6, 1, 2, 8, 5, 10}
	MergeSort1(s)
	fmt.Println(s)
	s1 := []int{3, 5, 4, 1, 2, 6, 1, 2, 8, 5, 10}
	fmt.Println(MergeSort2(s1))
}

func TestQuickSort(t *testing.T) {
	s := []int{3, 5, 4, 1, 2, 6, 1, 2, 8, 5, 10}
	QuickSort(s)
	fmt.Println(s)
}
