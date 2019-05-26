package stack_test

import (
	"testing"

	"github.com/jlevitt/katas/polish/stack"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStack(t *testing.T) {
	s := stack.NewStack()

	assert.Equal(t, 0, s.Size())
	_, ok := s.Pop()

	assert.Equal(t, false, ok)
}

func TestPushValues(t *testing.T) {
	s := stack.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Size())
}

func TestPushAndPop(t *testing.T) {
	s := stack.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())

	val, ok := s.Pop()

	assert.True(t, ok)
	assert.Equal(t, 3, val)
	assert.Equal(t, 2, s.Size())

	s.Push(4)
	assert.Equal(t, 3, s.Size())

	val, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 4, val)
	assert.Equal(t, 2, s.Size())

	val, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	assert.Equal(t, 1, s.Size())

	val, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 0, s.Size())

	_, ok = s.Pop()
	assert.False(t, ok)
}
