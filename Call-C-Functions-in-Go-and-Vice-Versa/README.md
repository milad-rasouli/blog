# Call C Functions in Go and Vice Versa

## Types

### String:	
```go
hostname := C.CString("100.100.24.81")
defer C.free(unsafe.Pointer(hostname))
```

### Int

```go
port := C.int(102)
```

### Struct
```c
struct Greetee {
    const char *name;
    int year;
};
```

```go
g := C.struct_Greetee{
    name: name,
    year: year,
}
```

## Print Hello World 
In this example we want to print something with the help of C function.
1. we need to  ``` import "C" ```.
2. we can write C code right above ``` import "C" ``` line. you must comment them out. it doesn't matter that they are commented out they will work fine.
3. import the C libraries that we need. Here we need ``` // #include <stdio.h> ``` and ``` // #include <stdlib.h> ```
4- now, all the function of the imported libraries, has added to "C" namespace. see the example to find out how it works. 

```go
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

```

Run it :
```bash
CGO_ENABELED=1 go build -o ./bin/hello ./main.go

./bin/hello
```

Result:
```
This is a "Hello World" message from C.
```

