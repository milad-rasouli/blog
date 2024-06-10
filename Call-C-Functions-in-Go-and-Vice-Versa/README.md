# Call C Functions in Go and Vice Versa

## Types

## String:	
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


