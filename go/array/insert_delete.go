package array

import (
	"errors"
)

// 无序数组，不考虑排序
func insert(old []int, number, pos int) (new []int, err error) {
	if pos < 0 {
		return nil, errors.New("index out of range")
	}
	new = append(old, old[pos])
	new[pos] = number
	return new, nil
}


