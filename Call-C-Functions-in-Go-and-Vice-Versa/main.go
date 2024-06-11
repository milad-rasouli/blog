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
