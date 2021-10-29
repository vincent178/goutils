package cursor

import (
	"math/rand"
	"testing"
)

func TestFileCursor_GetCursor(t *testing.T) {
	tests := []struct {
		name       string
		fileName   string
		cursorName string
		want       int
	}{
		{
			name:       "test GetCursor",
			fileName:   "fixtures/not_exist",
			cursorName: "test",
			want:       -1,
		},
		{
			name: "test GetCursor existed",
			fileName:   "fixtures/get_cursor1.txt",
			cursorName: "test",
			want:       1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewFileCursor(tt.fileName)
			if got := c.Get(tt.cursorName); got != tt.want {
				t.Errorf("FileCursor.GetCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileCursor_SaveCursor(t *testing.T) {
	tests := []struct {
		name    string
		fileName string
		cursorName string
		wantErr bool
	}{
		{
			name: "test SaveCursor",
			fileName: "fixtures/save_cursor1.txt",
			cursorName: "test",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewFileCursor(tt.fileName)
			n := rand.Intn(100)
			c.data = map[string]int{
				tt.cursorName: n,
			}

			if err := c.Save(); (err != nil) != tt.wantErr {
				t.Errorf("FileCursor.SaveCursor() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got := c.Get(tt.cursorName); got != n {
				t.Errorf("FileCursor.GetCursor() = %v, want %v", got, n)
			}
		})
	}
}

func TestFileCursor_UpdateCursor(t *testing.T) {
	tests := []struct {
		name    string
		fileName string
		cursorName string
	}{
		{
			name: "test SaveCursor",
			fileName: "fixtures/save_cursor1.txt",
			cursorName: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := rand.Intn(100)
			c, _ := NewFileCursor(tt.fileName)
			c.Update(tt.cursorName, n)
			if got := c.Get(tt.cursorName); got != n {
				t.Errorf("FileCursor.GetCursor() = %v, want %v", got, n)
			}
		})
	}
}
