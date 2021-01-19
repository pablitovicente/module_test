package bar

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Speak bars
func Speak(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)
	time.Sleep(time.Duration(n) * time.Second)

	fmt.Println("BAR...", id)
}
