package leet

/*
42. 接雨水
困难

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105
*/

func trap(height []int) int {
	length := len(height)

	leftMax := make([]int, length)
	leftMax[0] = height[0]

	rightMax := make([]int, length)
	rightMax[length-1] = height[length-1]

	for i := 1; i < length-1; i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	for i := length - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	area := 0
	for i := 1; i < length-1; i++ {
		if leftMax[i] > rightMax[i] {
			area += rightMax[i] - height[i]
		} else {
			area += leftMax[i] - height[i]
		}
	}

	return area
}

func trap1(height []int) int {
	length := len(height)
	leftMax := height[0]
	rightMax := height[length-1]

	area := 0
	i := 0
	j := length - 1

	for i < j {
		if leftMax < rightMax {
			area += leftMax - height[i]
			i++

			if leftMax < height[i] {
				leftMax = height[i]
			}
		} else {
			area += rightMax - height[j]
			j--

			if rightMax < height[j] {
				rightMax = height[j]
			}
		}
	}

	return area
}
