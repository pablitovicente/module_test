package main

import (
	"sync"

	"github.com/pablitovicente/module_test/bar"
	"github.com/pablitovicente/module_test/baz"
	"github.com/pablitovicente/module_test/foo"
)

func main() {
	Hallo()
}

// Hallo calls some methods using go routines
func Hallo() {
	var wg sync.WaitGroup

	wg.Add(3)

	go bar.Bar(&wg, 1)
	go baz.Baz(&wg, 2)
	go foo.Foo(&wg, 3)

	wg.Wait()
}
