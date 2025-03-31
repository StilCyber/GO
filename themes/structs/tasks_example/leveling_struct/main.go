package main

import (
	"fmt"
	"unsafe"
)

type data1 struct {
	aaa bool
	bbb int32 
	ccc bool
}

type data2 struct {
	aaa int32 
	bbb bool 
	ccc bool
}

func main() {
	fmt.Println(unsafe.Sizeof(data1{}))
	fmt.Println(unsafe.Sizeof(data2{}))
}
