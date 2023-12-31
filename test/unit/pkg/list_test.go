package pkg_test

import (
	"testing"

	"github.com/devcael/go-utils/pkg/list"
)

func Test_ShouldCreateList(t *testing.T) {
	l := list.New[int]()
	if l == nil {
		t.Error("List is nil")
	}
}

func Test_ShouldInsertItem(t *testing.T) {
	l := list.New[int]()
	l.Insert(1)
	if l.IsEmpty() {
		t.Error("List is empty")
	}
}

func Test_ShouldGetItem(t *testing.T) {
	l := list.New[int]()
	l.Insert(1)
	if l.Get(0) != 1 {
		t.Error("Item not found")
	}
}

func Test_ShouldSetItem(t *testing.T) {
	l := list.New[int]()
	l.Insert(1)
	l.Set(0, 2)
	if l.Get(0) != 2 {
		t.Error("Item not found")
	}
}

func Test_ShouldRemoveByItem(t *testing.T) {
	l := list.New[int]()
	l.Insert(1)
	l.RemoveByItem(1)
	if l.IsNotEmpty() {
		t.Error("Item not removed")
	}
}

func Test_ShouldRemoveItem(t *testing.T) {
	l := list.New[int]()
	l.Insert(1)
	l.Remove(0)
	if l.IsNotEmpty() {
		t.Error("Item not removed")
	}
}

func Test_ShouldInitializeListWithItemsAndRemoveAll(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	l.RemoveAll()
	if l.IsNotEmpty() {
		t.Error("List not empty")
	}
}

func Test_ShouldGetFirstItem(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	if l.First() != 1 {
		t.Error("First item not found")
	}
}

func Test_ShouldGetLastItem(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	if l.Last() != 3 {
		t.Error("Last item not found")
	}
}

func Test_ShouldDeleteItemByIndex(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	l.Remove(2)
	l.Remove(1)
	if l.Size() >= 2 {
		t.Error("Item not removed")
	}
}

func Test_ShouldGetIndexOfItem(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	if l.IndexOf(2) != 1 {
		t.Error("Index not found")
	}
}

func Test_ShouldNotGetIndexOfItem(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	if l.IndexOf(4) != -1 {
		t.Error("Index found")
	}
}

func Test_ShouldFindItem(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	if l.Find(2) != 2 {
		t.Error("Item not found")
	}
}

func Text_ShouldFindItemByIndex(t *testing.T) {
	l := list.FromList[int]([]int{1, 2, 3})
	if l.FindByIndex(1) != 2 {
		t.Error("Item not found")
	}
}
