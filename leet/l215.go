package leet

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：
1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func findKthLargest(nums []int, k int) int {
	head := &node{
		val: nums[0],
	}

	length := len(nums)
	asc := true
	if k > length/2 {
		k = length - k + 1
		asc = false
	}

	for i := 1; i < len(nums); i++ {
		if asc {
			head = insertAsc(head, nums[i])
		} else {
			head = insertDesc(head, nums[i])
		}

		if i > k-1 {
			head = head.next
		}
	}
	return head.val
}

type node struct {
	val  int
	next *node
}

func insertAsc(head *node, val int) *node {
	n := &node{
		val: val,
	}
	cur := head

	if val <= head.val {
		n.next = head
		return n
	}

	for cur != nil {
		if cur.next == nil {
			cur.next = n
			break
		} else {
			if val > cur.next.val {
				cur = cur.next
			} else {
				n.next = cur.next
				cur.next = n
				break
			}
		}
	}

	return head
}

func insertDesc(head *node, val int) *node {
	n := &node{
		val: val,
	}
	cur := head

	if val >= head.val {
		n.next = head
		return n
	}

	for cur != nil {
		if cur.next == nil {
			cur.next = n
			break
		} else {
			if val < cur.next.val {
				cur = cur.next
			} else {
				n.next = cur.next
				cur.next = n
				break
			}
		}
	}

	return head
}
