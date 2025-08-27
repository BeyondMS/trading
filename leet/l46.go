package leet

/*
46. 全排列

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

func permute(nums []int) [][]int {
	var res [][]int
	dfs(&res, nums, 0)
	return res
}

func dfs(res *[][]int, nums []int, k int) {
	if k == len(nums)-1 {
		*res = append(*res, append([]int{}, nums...))
		return
	}

	for i := k; i < len(nums); i++ {
		nums[i], nums[k] = nums[k], nums[i]
		dfs(res, nums, k+1)
		nums[k], nums[i] = nums[i], nums[k]
	}
}

func permute1(nums []int) [][]int {
	var res [][]int
	var dfx func(x int)
	dfx = func(x int) {
		if x == len(nums)-1 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := x; i < len(nums); i++ {
			nums[i], nums[x] = nums[x], nums[i]
			dfx(x + 1)
			nums[i], nums[x] = nums[x], nums[i]
		}
	}
	dfx(0)
	return res
}

func permute2(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var res [][]int
	for i := 0; i < len(nums); i++ {
		var subArr []int
		for j := 0; j < len(nums); j++ {
			if i != j {
				subArr = append(subArr, nums[j])
			}
		}
		subRes := permute(subArr)

		for j := 0; j < len(subRes); j++ {
			r := append(subRes[j], nums[i])
			res = append(res, r)
		}
	}
	return res
}
