package main

import (
	"context"
	"time"
)

func Query(string) string

func DoQuery(queryStr string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resultCh := make(chan string, 1)
	go func() {
		result := Query(queryStr)
		resultCh <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-resultCh:
		return result, nil
	}
}

func main() {

}
