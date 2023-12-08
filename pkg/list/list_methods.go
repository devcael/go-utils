package list

func (l *List[T]) Each(fn func(item T)) {
	for _, i := range l.data {
		fn(i)
	}
}

func (l *List[T]) RemoveWhere(fn func(item T) bool) {
	for i, item := range l.data {
		if fn(item) {
			l.data = append(l.data[:i], l.data[i+1:]...)
		}
	}
}

func FilterImutable() {

}

func Filter() {

}
