package goutils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	Name      string `csv:"Name"`
	Age       uint   `csv:"Age"`
	Height    int    `csv:"Height"`
	IsTeacher bool   `csv:"Is Teacher"`
}

var pi interface{} = &person{}

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
					"Name":       "Jojo",
					"Age":        "22",
					"Height":     "188",
					"Is Teacher": "false",
				},
				out: &person{},
			},
		},
		{
			name: "with empty value",
			args: args{
				src: map[string]string{
					"Name":       "Jojo",
					"Age":        "",
					"Height":     "188",
					"Is Teacher": "false",
				},
				out: &person{},
			},
		},
		{
			name: "with interface out",
			args: args{
				src: map[string]string{
					"Name":       "Jojo",
					"Age":        "",
					"Height":     "188",
					"Is Teacher": "false",
				},
				out: pi,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CsvMapToStruct(tt.args.src, tt.args.out)
			assert.NoError(t, err)

			var age int
			if tt.args.src["Age"] != "" {
				age, _ = strconv.Atoi(tt.args.src["Age"])
			}

			p := tt.args.out.(*person)
			assert.Equal(t, p.Name, "Jojo")
			assert.Equal(t, p.Age, uint(age))
			assert.Equal(t, p.Height, 188)
			assert.Equal(t, p.IsTeacher, false)
		})
	}
}
