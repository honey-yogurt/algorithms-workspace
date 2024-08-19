package array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	old := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newer, err := insert(old, 10, 5)
	if err != nil {
		return
	}
	fmt.Println(newer)
}
