package goutils

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// make sure Error implement error interface
var _ error = &Error{}

type Error struct {
	Err   error
	Stack []string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s\n%s\n", e.Err.Error(), strings.Join(e.Stack, "\n"))
}

func NewError(text string) *Error {
	stacks := make([]string, 0)
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	for true {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))

		if !more {
			break
		}
	}

	return &Error{
		Err:   errors.New(text),
		Stack: stacks,
	}
}
