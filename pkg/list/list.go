package list

import (
	"errors"

	. "github.com/devcael/go-utils/pkg/str"
)

type List[T comparable] struct {
	data []T
}

func (l *List[T]) Get(index int) T {
	return l.data[index]
}

func (l *List[T]) Set(index int, value T) {
	if index >= l.Size() {
		panic("Index out of bound")
	}
	l.data[index] = value
}

func (l *List[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *List[T]) IsNotEmpty() bool {
	return !l.IsEmpty()
}

func (l *List[T]) Last() T {
	if len(l.data) == 0 {
		panic("Lista vazia")
	}
	return l.data[len(l.data)-1]
}

func (l *List[T]) First() T {
	if len(l.data) == 0 {
		panic("Lista vazia")
	}
	return l.data[0]
}

func (l *List[T]) GetItem(item T) T {
	index := l.IndexOf(item)
	if index == -1 {
		panic("Valor não encontrado")
	}
	return l.data[index]
}

func (l *List[T]) validateIndexOutBound(index int) error {
	if len(l.data) < index {

		message := Concat([]string{
			"Index out of bound: ",
			ToString(index),
		})

		return errors.New(message)
	}
	return nil
}

func FromList[T comparable](values []T) *List[T] {
	return &List[T]{
		data: values,
	}
}

func New[T comparable]() *List[T] {
	return &List[T]{
		data: []T{},
	}
}

func (l *List[T]) Size() int {
	return len(l.data)
}

func (l *List[T]) Insert(item T) *List[T] {
	l.data = append(l.data, item)
	return l
}

func (l *List[T]) IndexOf(i T) int {
	for it := 0; it < len(l.data); it++ {
		if l.data[it] == i {
			return it
		}
	}
	return -1
}

func (l *List[T]) Find(item T) T {
	for it := 0; it < l.Size(); it++ {
		if l.data[it] == item {
			return l.data[it]
		}
	}
	panic("Valor não encontrado")
}

func (l *List[T]) FindByIndex(i int) T {
	err := l.validateIndexOutBound(i)
	if err != nil {
		panic(err)
	}

	for it := 0; it < l.Size(); it++ {
		if i == it {
			return l.data[it]
		}
	}
	panic("Valor não encontrado")
}
