package leet

/*
84. 柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:

输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

示例 2：

输入： heights = [2,4]
输出： 4


提示：

1 <= heights.length <=105
0 <= heights[i] <= 104
*/

func largestRectangleArea(heights []int) int {
	var leftStack []int
	var rightStack []int

	length := len(heights)
	leftResult := make([]int, length)
	rightResult := make([]int, length)

	for i := 0; i < length; i++ {
		tmp := -1

		if len(leftStack) == 0 {
			leftResult[i] = -1
			leftStack = []int{i}
			continue
		}

		for j := len(leftStack) - 1; j >= 0; j-- {
			if heights[i] > heights[leftStack[j]] {
				tmp = j
				break
			}
		}
		if tmp == -1 {
			leftResult[i] = -1
			leftStack = []int{i}
		} else {
			leftResult[i] = leftStack[tmp]
			leftStack = leftStack[:tmp+1]
			leftStack = append(leftStack, i)
		}
	}

	for i := length - 1; i >= 0; i-- {
		if len(rightStack) == 0 {
			rightResult[i] = length
			rightStack = []int{i}
			continue
		}

		tmp := length
		for j := len(rightStack) - 1; j >= 0; j-- {
			if heights[i] > heights[rightStack[j]] {
				tmp = j
				break
			}
		}
		if tmp == length {
			rightResult[i] = length
			rightStack = []int{i}
		} else {
			rightResult[i] = rightStack[tmp]
			rightStack = rightStack[:tmp+1]
			rightStack = append(rightStack, i)
		}
	}

	maxArea := 0
	for i := 0; i < length; i++ {
		area := (rightResult[i] - leftResult[i] - 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// timeout
func largestRectangleArea1(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		area := heights[i]
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				break
			}
			area += heights[i]
		}
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < heights[i] {
				break
			}
			area += heights[i]
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
