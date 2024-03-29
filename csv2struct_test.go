package goutils

import (
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
		src           map[string]string
		suppressError bool
		caseInsensive bool
		out           interface{}
	}
	tests := []struct {
		name string
		args args
		want *person
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
			want: &person{
				Name:      "Jojo",
				Age:       22,
				Height:    188,
				IsTeacher: false,
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
			want: &person{
				Name:      "Jojo",
				Age:       0,
				Height:    188,
				IsTeacher: false,
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
			want: &person{
				Name:      "Jojo",
				Age:       0,
				Height:    188,
				IsTeacher: false,
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
				suppressError: true,
				out:           &person{},
			},
			want: &person{
				Name:   "Jojo",
				Age:    0,
				Height: 0,
			},
		},
		{
			name: "with caseInsensive true",
			args: args{
				src: map[string]string{
					"Name":       "Jojo",
					"age":        "",
					"height":     "188",
					"is teacher": "true",
				},
				caseInsensive: true,
				out:           &person{},
			},
			want: &person{
				Name:      "Jojo",
				Age:       0,
				Height:    188,
				IsTeacher: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.args.suppressError {
				err = CsvMapToStruct(tt.args.src, tt.args.out, WithSuppressError(true))

			} else if tt.args.caseInsensive {
				err = CsvMapToStruct(tt.args.src, tt.args.out, WithCaseInsensitive(true))
			} else {
				err = CsvMapToStruct(tt.args.src, tt.args.out)
			}

			assert.NoError(t, err)

			p := tt.args.out.(*person)
			assert.Equal(t, p, tt.want)
		})
	}
}
