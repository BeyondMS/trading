package leet

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"c1", args{"3[a2[c]]"}, "accaccacc"},
		{"c2", args{"a3[b2[c]d]e"}, "abccdbccdbccde"},
		{"c3", args{"3[a2[c]]"}, "accaccacc"},
		{"c4", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"c5", args{"abc3[cd]xyz"}, "abccdcdcdxyz"},
		{"c6", args{"a"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
