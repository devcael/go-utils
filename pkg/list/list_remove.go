package list

func (l *List[T]) RemoveByItem(item T) *List[T] {
	for i := 0; i < l.Size(); i++ {
		currItem := l.data[i]
		if currItem == item {
			l.data = append(l.data[:i], l.data[i+1:]...)
			return l
		}
	}

	panic("Index outbound cannot remove")
}

func (l *List[T]) Remove(index int) *List[T] {
	err := l.validateIndexOutBound(index)
	if err != nil {
		panic(err)
	}
	for i := 0; i < l.Size(); i++ {
		if i == index {
			l.data = append(l.data[:i], l.data[i+1:]...)
			return l
		}
	}

	panic("Index outbound cannot remove")
}

func (l *List[T]) RemoveAll() *List[T] {
	l.data = []T{}
	return l
}
