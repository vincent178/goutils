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
		name          string
		args          args
		suppressError bool
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
		{
			name: "with invalid value error",
			args: args{
				src: map[string]string{
					"Name":       "Jojo",
					"Age":        "",
					"Height":     "L",
					"Is Teacher": "false",
				},
				out: pi,
			},
			suppressError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.suppressError {
				err = CsvMapToStruct(tt.args.src, tt.args.out, WithSuppressError(true))
			} else {
				err = CsvMapToStruct(tt.args.src, tt.args.out)
			}

			assert.NoError(t, err)

			var age int
			if tt.args.src["Age"] != "" {
				age, _ = strconv.Atoi(tt.args.src["Age"])
			}

			height := 188
			if tt.args.src["Height"] == "L" {
				height = 0
			}

			p := tt.args.out.(*person)
			assert.Equal(t, p.Name, "Jojo")
			assert.Equal(t, p.Age, uint(age))
			assert.Equal(t, p.Height, height)
			assert.Equal(t, p.IsTeacher, false)
		})
	}
}
