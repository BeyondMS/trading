package leet

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"c1", args{[]int{1, 2, 3}}, [][]int{{3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}, {2, 1, 3}, {1, 2, 3}}},
		{"c2", args{[]int{1, 2}}, [][]int{{2, 1}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
