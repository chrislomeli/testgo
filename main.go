package main
/*
#include <stdio.h>
char * helloc() { return "Hello from C"; }
 */
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.helloc()))
}
