//学习切片的本质结构SliceHeader。

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Slice []int

func (A Slice) Append(value int) {
	A1 := append(A, value)
	fmt.Printf("%p\n%p\n", A, A1)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&A))
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n", sh.Data, sh.Len, sh.Cap)

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&A1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n", sh1.Data, sh1.Len, sh1.Cap)
}

func main() {
	mSlice := make(Slice, 10, 20)
	mSlice.Append(5)
	fmt.Println(mSlice)
}
