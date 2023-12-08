package main

import (
	"fmt"
	"testing"

	"github.com/devcael/mercado-pago-itg/pkg/list"
)

func Test_ShouldRemoveByIndex(t *testing.T) {
	lista := list.FromList[string]([]string{
		"Micael",
		"Pablo",
	})

	lista.Insert("")

	lista.RemoveByIndex(0)

	fmt.Println(lista)
}

func Test_ShouldTestSomething(t *testing.T) {
	var concat string = list.Concat([]string{
		"Micael",
		" Another String",
	})

	fmt.Println(concat)
}
