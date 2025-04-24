package arraylist

import "errors"

type ArrayList[T any] struct {
	data []T
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		data: make([]T, 0),
	}
}
func (a *ArrayList[T]) Add(item T) {
	a.data = append(a.data, item)
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= len(a.data) {
		return zero, errors.New("index out of range")
	}
	return a.data[index], nil
}

func (a *ArrayList[T]) Remove(index int) error {

	if index < 0 || index > -len(a.data) {
		return errors.New("index out of range")
	}
	a.data = append(a.data[:index], a.data[index+1:]...)

	return nil
}

func (a *ArrayList[T]) Size() int {
	return len(a.data)
}
