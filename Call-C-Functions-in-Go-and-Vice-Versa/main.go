package main

/*
#include <stdio.h>
#include <stdlib.h>

struct Greetee {
    const char *name;
    int year;
};

void sayHello() {
    printf("Hello from C code\n");
}

int greet(const char *name, int year, char *out) {
    int n;
    n = sprintf(out, "Greetings, %s from %d! We come in peace :)", name, year);
    return n;
}

int greetWithStruct(struct Greetee *g, char *out) {
    int n;
    n = sprintf(out, "Greetings, %s from %d! We come in peace :)", g->name, g->year);
    return n;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	{
		C.sayHello()
	}

	{
		name := C.CString("Mili")
		defer C.free(unsafe.Pointer(name))
		year := C.int(2023)
		ptr := C.malloc(C.sizeof_char * 1024)
		defer C.free(unsafe.Pointer(ptr))

		size := C.greet(name, year, (*C.char)(ptr))
		msg := C.GoBytes(ptr, size)
		fmt.Printf("%s\n", string(msg))
	}

	{
		name := C.CString("Mili Mili")
		defer C.free(unsafe.Pointer(name))

		year := C.int(2024)
		g := C.struct_Greetee{
			name: name,
			year: year,
		}

		ptr := C.malloc(C.sizeof_char * 1024)
		defer C.free(unsafe.Pointer(ptr))
		size := C.greetWithStruct(&g, (*C.char)(ptr))
		msg := C.GoBytes(ptr, size)
		fmt.Println(string(msg))
	}
}
