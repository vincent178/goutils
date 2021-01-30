package goutils

import (
	"fmt"
	"testing"
)

func testFunc() {
	NewError("hello world")

}

func TestNewError(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *Error
	}{
		{
			name: "test stack",
			args: "hello world",
			want: &Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewError(tt.args)
			fmt.Println(got)
		})
	}
}
