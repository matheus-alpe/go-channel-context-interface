package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func printSize[T any](value T) {
	fmt.Println(reflect.TypeOf(value), "size:", unsafe.Sizeof(value), "| bits:", unsafe.Sizeof(value)*8)
}

func Example01() {
	var a bool
	printSize(a)

	var b struct{}
	printSize(b)

	var d int
	printSize(d)
	var du uint
	printSize(du)

	var e int8
	printSize(e)
	var eu uint8
	printSize(eu)

	var f int16
	printSize(f)
	var fu uint16
	printSize(fu)

	var g int32
	printSize(g)
	var gu uint32
	printSize(gu)

	var h int64
	printSize(h)
	var hu uint64
	printSize(hu)
}
