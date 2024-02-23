package main

import (
	"unsafe"

	"github.com/zhangga/swagger/hotfix/example/originfunc"
)

func main() {
	println(originfunc.GetSpecialName(1))

	//ch := make(chan struct{})
	//ch <- struct{}{}

	buf := new([]byte)
	copy(*buf, []byte("hello"))
	size := unsafe.Sizeof(htest{})
	fsize := uintptr(-int(size))
	ffsize := fsize & 0x7
	println(size + ffsize)
}

type htest struct {
	age uint32
}
