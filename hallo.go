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

	go bar.Bar(&wg)
	go baz.Baz(&wg)
	go foo.Foo(&wg)

	wg.Wait()
}
