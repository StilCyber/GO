package main

import (
	"log"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			value := 0
			for j := 0; j < 10; j++ {
				mutex.Lock()
				value++
				mutex.Unlock()
			}

			log.Println(value)
		}()
	}

	wg.Wait()
}