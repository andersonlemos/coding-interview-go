package arraylist

import "errors"

type ArrayList struct {
	data []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0),
	}
}
func (a *ArrayList) Add(item interface{}) {
	a.data = append(a.data, item)
}

func (a *ArrayList) Size() int {
	return len(a.data)
}

func (a *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(a.data) {
		return nil, errors.New("index out of range")
	}
	return a.data[index], nil
}
