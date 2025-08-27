package leet

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"case2", args{[]int{3, 2, 1, 5, 6, 4}, 3}, 4},
		{"case3", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
		{"case4", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3}, 5},
		{"case5", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 9}, 1},
		{"case6", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 8}, 2},
		{"case7", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 7}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
