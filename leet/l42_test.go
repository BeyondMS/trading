package leet

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"c1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0},
		{"c2", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"c3", args{[]int{4, 2, 0, 3, 2, 5}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"c1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 0},
		{"c2", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"c3", args{[]int{4, 2, 0, 3, 2, 5}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap1(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}
