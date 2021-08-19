package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	Name string `csv:"Name"`
	Age  int    `csv:"Age"`
}

func TestMapToStruct(t *testing.T) {
	type args struct {
		src map[string]string
		out interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "unmarshal to struct",
			args: args{
				src: map[string]string{
					"Name": "Jojo",
					"Age":  "100",
				},
				out: &person{},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CsvMapToStruct(tt.args.src, tt.args.out)
			p := tt.args.out.(*person)
			assert.Equal(t, p.Name, "Jojo")
			assert.Equal(t, p.Age, 100)
		})
	}
}
