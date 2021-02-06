package goutils

import (
	"reflect"
	"testing"
	"time"
)

var layout = "2006-01-02 15:04:05"

func TestStartOfDay(t *testing.T) {
	arg, _ := time.ParseInLocation(layout, "2016-06-08 21:16:15", time.Local)
	want, _ := time.ParseInLocation(layout, "2016-06-08 00:00:00", time.Local)

	tests := []struct {
		name string
		args time.Time
		want time.Time
	}{

		{
			name: "start of day",
			args: arg,
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfDay(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfMonth(t *testing.T) {
	arg, _ := time.ParseInLocation(layout, "2016-06-08 21:16:15", time.Local)
	want, _ := time.ParseInLocation(layout, "2016-06-01 00:00:00", time.Local)

	tests := []struct {
		name string
		args time.Time
		want time.Time
	}{
		{
			name: "start of month",
			args: arg,
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfMonth(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfYear(t *testing.T) {
	arg, _ := time.ParseInLocation(layout, "2016-06-08 21:16:15", time.Local)
	want, _ := time.ParseInLocation(layout, "2016-01-01 00:00:00", time.Local)

	tests := []struct {
		name string
		args time.Time
		want time.Time
	}{
		{
			name: "start of year",
			args: arg,
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfYear(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
