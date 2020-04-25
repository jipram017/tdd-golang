package stack

import "fmt"

var ErrNoSuchElement = fmt.Errorf("No such element")

func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

type Stack struct {
	arr []int
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(value int) error {
	s.arr = append(s.arr, value)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	x, lastIdx := s.GetLastElement()
	s.arr = s.arr[:lastIdx]
	return x, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	x, _ := s.GetLastElement()
	return x, nil
}

func (s *Stack) GetLastElement() (int, int) {
	lastIdx := s.Size() - 1
	val := s.arr[lastIdx]
	return val, lastIdx
}
