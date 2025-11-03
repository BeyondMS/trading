package leet

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"c1", args{[]int{1, 2, 3}, 3}, 2},
		{"c2", args{[]int{1, 1, 1}, 2}, 2},
		{"c2", args{[]int{-1, 1, 1, -2, 1, -1, -1}, 0}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
