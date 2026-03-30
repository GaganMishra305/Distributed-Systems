package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(42);

	count := 0
	finished := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()

			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
		} ()
	}


	for {
		mu.Lock()
		if count >= 5 || finished == 10 {
			break
		}
		mu.Unlock()
	}
}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(100)))
	return rand.Intn(10) < 5
}

// we can do the same thing above using a channel