package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	cs := C.CString("This is a \"Hello World\" message from C.\n")
	defer func() {
		C.free(unsafe.Pointer(cs))
	}()
	C.puts(cs)
}
