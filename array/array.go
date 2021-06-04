package main

//#cgo CFLAGS:
//#cgo LDFLAGS: -libverbs
//#include <infiniband/verbs.h>
//#include <stdio.h>
/*
void write(uint8_t *gid) {
	uint8_t a[16];
	int m;
	for (m =0; m<16; m++) {
		a[m] = m;
	}
	memcpy(gid, a, 16);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	var gid [16]uint8
	fmt.Println(&gid)
	C.write((*C.uint8_t)(unsafe.Pointer(&gid[0])))
	for i := 1; i < 16; i++ {
		fmt.Println(&gid[i], "\t")
	}
	fmt.Println(&gid)
}
