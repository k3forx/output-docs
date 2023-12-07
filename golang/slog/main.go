package main

import (
	"fmt"
	"unsafe"
)

type stringptr *byte

type Value struct {
	_   [0]func()
	num uint64
	any any
}

func StringValue(value string) Value {
	return Value{num: uint64(len(value)), any: stringptr(unsafe.StringData(value))}
}

func main() {
	str := "abc"

	var s any = str
	fmt.Println(s)

	var ss Value = StringValue(str)
	fmt.Println(ss)

	fmt.Printf("byte size of str: %d\n", unsafe.Sizeof(str))
	fmt.Printf("byte size of s: %d\n", unsafe.Sizeof(s))
	fmt.Printf("byte size of ss: %d\n", unsafe.Sizeof(ss))
}
