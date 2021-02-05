package goutils

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "merge string slice",
			args: args{
				s1: []string{"a", "b"},
				s2: []string{"c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "s1 is nil",
			args: args{
				s1: nil,
				s2: []string{"c", "d"},
			},
			want: []string{"c", "d"},
		},
		{
			name: "s1 is empty",
			args: args{
				s1: []string{},
				s2: []string{"c", "d"},
			},
			want: []string{"c", "d"},
		},
		{
			name: "s2 is nil",
			args: args{
				s1: []string{"a", "b"},
				s2: nil,
			},
			want: []string{"a", "b"},
		},
		{
			name: "s2 is empty",
			args: args{
				s1: []string{"a", "b"},
				s2: []string{},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
