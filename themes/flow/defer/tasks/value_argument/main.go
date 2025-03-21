package main

import "fmt"

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func notify(status string) {
	fmt.Println(status)
}

func process() {
	var status string
	defer func(s string) {
		notify(s)
	}(status)

	status = StatusError
}

func main() {
	process()
}
