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

### Run it :
```bash
$ go env -w CGO_ENABLED=1
$ go build -o ./bin/hello ./main.go

$ ./bin/hello
```
> [!NOTE]
> To build the program you have to have GCC or G++ on your machine. [Linux](https://www.cherryservers.com/blog/how-to-install-gcc-on-ubuntu), [Windows](https://jmeubank.github.io/tdm-gcc/download/) *Don't forget to check add to path while installing*

### Result:
```
This is a "Hello World" message from C.
```

## Call Your C Function in Go
In this example I will show you how you can write C function with different arguments and use them in your Go program


```go

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

```


### Run it :
```bash
$ go env -w CGO_ENABLED=1
$ go build -o ./bin/hello ./main.go

$ ./bin/hello
```

### Result:
```
Hello from C code
Greetings, Mili from 2023! We come in peace :)
Greetings, Mili Mili from 2024! We come in peace :)
```

## Call Go Function in C
Here you can see how to call a Go function in a C function.
1. put ``` //export GoPrintNameAge ``` above your go function.
2. put ``` extern void GoPrintNameAge(char*,int); ``` in your c code, then you can call it.
3. build it with this flag ``` -ldflags='-extldflags=-Wl,--allow-multiple-definition' ```

```go 
package main

/*
#include <stdio.h>

extern void GoPrintNameAge(char*,int);

void getNameAge(){
	char name[100];
	int age;

	printf("Enter Your Name: ");
	scanf("%s",&name);

	printf("Enter Your Age: ");
	scanf("%d",&age);

	GoPrintNameAge(name,age);
}
*/
import "C"
import (
	"fmt"
)

//export GoPrintNameAge
func GoPrintNameAge(name *C.char, age C.int) {
	goName := C.GoString(name)
	goAge := int(age)
	fmt.Printf("Hello, %s!\nYour age is %d\n", goName, goAge)
}

func main() {
	C.getNameAge()
}

```

### Run it:
```bash
$ go env -w CGO_ENABLED=1
$ go build  -ldflags='-extldflags=-Wl,--allow-multiple-definition' -o ./bin/hello ./main.go

$ ./bin/hello
```

### Result:

```bash

Enter Your Name: Milad
Enter Your Age: 10
Hello, Milad!
Your age is 10

```