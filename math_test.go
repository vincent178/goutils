package goutils

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		nums []float64
		want float64
	}{
		{
			name: "does not pass numbers",
			nums: []float64{},
			want: 0,
		},
		{
			name: "1 number",
			nums: []float64{0.1},
			want: 0.1,
		},
		{
			name: "2 number",
			nums: []float64{0.1, 0.2},
			want: 0.2,
		},
		{
			name: "3 number",
			nums: []float64{0.1, 0.3, 0.2},
			want: 0.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.nums...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
