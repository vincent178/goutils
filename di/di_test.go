package di

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestService struct {
	init bool
	closed bool
}

func (s *TestService) Init() error {
	s.init = true
	return nil
}

func (s *TestService) Close() error {
	s.closed = true
	return nil
}

func TestLoad(t *testing.T) {
	var ts TestService
	err := Load(&ts)
	assert.NoError(t, err)
	assert.True(t, ts.init)
}

func TestClose(t *testing.T) {
	var ts TestService
	container.services["TestService"] = &ts

	err := Close()
	assert.NoError(t, err)
	assert.True(t, ts.closed)
}