package main

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: -L./lib -lhello
#include "./lib/hello.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.init_stuff()
	//C.hello(C.CString("John Smith"))
	src := []byte{1, 2, 3}
	dst := make([]byte, 3)

	srcPtr := unsafe.Pointer(&src[0])
	dstPtr := unsafe.Pointer(&dst[0])

	C.square_u8_array((*C.char)(srcPtr), 3, (*C.char)(dstPtr))
	fmt.Println(dst)
}
