code space of algo.

## 数组：
+ [两数之和](https://leetcode.cn/problems/two-sum/description/)
+ [三数之和](https://leetcode.cn/problems/3sum/description/)
+ [反转字符串](https://leetcode.cn/problems/reverse-string/description/)
+ [最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/description/)
+ [删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/)
+ [移除元素](https://leetcode.cn/problems/remove-element/description/)
+ [移动零](https://leetcode.cn/problems/move-zeroes/description/)





### 滑动窗口O(n)：

滑动窗口算法技巧主要用来解决**子数组问题**，比如让你**寻找符合某个条件的最长/最短子数组**。

滑动窗口算法技巧的思路就是**维护一个窗口，不断滑动，然后更新答案**，该算法的大致逻辑如下：

```go
left , right := 0,0
for right < len(nums) {
	// 增大窗口
	windows.addLast(nums[right])
	right++
	
	// 判断窗口是否需要收缩
	for windows need shrink {
		// 收缩窗口
		windows.removeFirst(nums[left])
		left++
    }
}
```

> 指针 left, right 不会回退（它们的值只增不减），所以字符串/数组中的每个元素都只会进入窗口一次，然后被移出窗口一次，不会说有某些元素多次进入和离开窗口，所以算法的时间复杂度就和字符串/数组的长度成正比。

```go
// 滑动窗口算法伪码框架
func slidingWindow(s string) {
    // 用合适的数据结构记录窗口中的数据，根据具体场景变通
    // 比如说，我想记录窗口中元素出现的次数，就用 map
    // 如果我想记录窗口中的元素和，就可以只用一个 int
    var window = ...

    left, right := 0, 0
    for right < len(s) {
        // c 是将移入窗口的字符
        c := rune(s[right])
        window[c]++
        // 增大窗口
        right++
        // 进行窗口内数据的一系列更新
        ...

        // *** debug 输出的位置 ***
        // 注意在最终的解法代码中不要 print
        // 因为 IO 操作很耗时，可能导致超时
        fmt.Println("window: [",left,", ",right,")")
        // ***********************

        // 判断左侧窗口是否要收缩
        for left < right && window needs shrink { //replace "window needs shrink" with actual condition
            // d 是将移出窗口的字符
            d := rune(s[left])
            window[d]--
            // 缩小窗口
            left++
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```

**框架中两处 ... 表示的更新窗口数据的地方，在具体的题目中，你需要做的就是往这里面填代码逻辑**。而且，这两个 ... 处的操作分别是扩大和缩小窗口的更新操作，等会你会发现它们操作是完全对称的。


算法技巧：
1、我们在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」。

> 为什么要「左闭右开」区间
>
> 理论上你可以设计两端都开或者两端都闭的区间，但设计为左闭右开区间是最方便处理的。
>
> 因为这样初始化 left = right = 0 时区间 [0, 0) 中没有元素，但只要让 right 向右移动（扩大）一位，区间 [0, 1) 就包含一个元素 0 了。
> 
> 如果你设置为两端都开的区间，那么让 right 向右移动一位后开区间 (0, 1) 仍然没有元素；如果你设置为两端都闭的区间，那么初始区间 [0, 0] 就包含了一个元素。这两种情况都会给边界处理带来不必要的麻烦。

2、我们先不断地增加 right 指针扩大窗口 [left, right)，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。

3、此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right)，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮结果。

4、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。

**第 2 步相当于在寻找一个「可行解」，然后第 3 步在优化这个「可行解」，最终找到最优解**

+ [无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
+ [字符串排列](https://leetcode.cn/problems/permutation-in-string/description/)
+ [最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/description/)
+ [找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/)


## 链表

+ [合并两个有序数组](https://leetcode.cn/problems/merge-two-sorted-lists/description/)
+ [分隔链表](https://leetcode.cn/problems/partition-list/description/)