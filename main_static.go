package main

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is necessary to fix the linker errors about `dlsym` that would otherwise appear.

/*
#cgo LDFLAGS: ./lib/libhello.a -ldl
#include "./lib/hello.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	C.init_stuff()
	src := []byte{1, 2, 3}
	dst := make([]byte, 3)

	srcPtr := unsafe.Pointer(&src[0])
	dstPtr := unsafe.Pointer(&dst[0])

	C.square_u8_array((*C.char)(srcPtr), 3, (*C.char)(dstPtr))
	fmt.Println(dst)
}
