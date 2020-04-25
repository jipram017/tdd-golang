package stack_test

import (
	"stack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := stack.New()
	t.Run("A stack is empty on construction", func(t *testing.T) {
		assert.True(t, s.IsEmpty())
	})

	t.Run("A stack has size 0 on construction", func(t *testing.T) {
		assert.Equal(t, 0, s.Size())
	})
}

func TestInsert(t *testing.T) {
	t.Run("After n items pushed to an empty stack (n>0), the stack is non-empty and the size is equals to n", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())
	})
}

func TestInsertAndGet(t *testing.T) {
	t.Run("If one pushes an element X then pop, the popped value is X and the size of stack is one less than old size", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(6)
		assert.Equal(t, 3, s.Size())
		x, _ := s.Pop()
		assert.Equal(t, x, 6)
		assert.Equal(t, s.Size(), 2)
	})

	t.Run("If one pushes an element X then peek, the peeked value is X but the size of the stack is still the same", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(6)
		assert.Equal(t, 3, s.Size())
		x, _ := s.Peek()
		assert.Equal(t, x, 6)
		assert.Equal(t, s.Size(), 3)
	})
}

func TestError(t *testing.T) {
	t.Run("Popping from an empty stack and retun error: NoSuchElement", func(t *testing.T) {
		s := stack.New()
		_, err := s.Pop()
		if err == nil {
			t.Fail()
			t.Logf("Expect error is not nil, but got %s", err)
		}
		assert.Equal(t, stack.ErrNoSuchElement, err)
	})

	t.Run("Peekeng from an empty stack and retun error: NoSuchElement", func(t *testing.T) {
		s := stack.New()
		_, err := s.Peek()
		if err == nil {
			t.Fail()
			t.Logf("Expect error is not nil, but got %s", err)
		}
		assert.Equal(t, stack.ErrNoSuchElement, err)
	})
}
