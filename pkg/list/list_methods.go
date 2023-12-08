package list

func (l *List[T]) Each(fn func(item T)) {
	for _, i := range l.data {
		fn(i)
	}
}

func (l *List[T]) WhereImutable(fn func(item T) bool) *List[T] {
	bufferList := New[T]()
	for _, item := range l.data {
		canAdd := fn(item)
		if canAdd {
			bufferList.Insert(item)
		}
	}
	return bufferList
}
