package pkg_test

import (
	"testing"

	"github.com/devcael/go-utils/pkg/list"
)

func TestSomething(t *testing.T) {
	lista := list.New[string]()
	lista2 := list.FromList[string]([]string{
		"Micael",
		"Micael",
		"Micael",
		"Micael",
	})

	lista2.Insert("")
	lista2.RemoveByItem("Micael")
	lista.RemoveByIndex(1)

	lista.Insert("Micael").Insert("Sla")

}
