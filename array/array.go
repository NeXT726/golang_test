package main

//#cgo CFLAGS:
//#cgo LDFLAGS: -libverbs
//#include <infiniband/verbs.h>
//#include <stdio.h>
/*
void write(uint8_t **gid) {
	uint8_t a[16];
	int m;
	for (m =1; m<16; m++) {
		a[m] = m;
	}
	memcpy(*gid, a, 16);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	var gid [16]uint8
	C.write((**C.uint8_t)(unsafe.Pointer(&gid)))
	fmt.Println(gid)
}
