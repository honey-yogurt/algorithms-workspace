package array

import (
	"fmt"
	"testing"
)

func TestMinWindows(t *testing.T) {
	var s = "cabwefgewcwaefgcf"
	var t1 = "cae"
	fmt.Println(minWindow(s, t1))
}

func TestMinWindows2(t *testing.T) {
	var s = "ADOBECODEBANC"
	var t1 = "ABC"
	fmt.Println(minWindow(s, t1))
}
