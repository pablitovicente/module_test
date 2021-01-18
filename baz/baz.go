package baz

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Baz does bars
func Baz(wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	time.Sleep(time.Duration(n) * time.Second)

	fmt.Println("...BAZ...")
}
