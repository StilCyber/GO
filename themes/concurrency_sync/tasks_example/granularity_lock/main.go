package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var cache map[string]string

func main() {
	mutex.Lock()
	item := cache["key"]
	fmt.Println(item)
	mutex.Unlock()
}
