//学习golang多维数组。多维数组的直接赋值。

package main

import (
	"fmt"
)

func main() {
	var a [5][9]int
	var b [5]int
	var c [9]int
	for i := 0; i < 45; i++ {
		a[i/9][i%9] = i
	}

	//b = a[1]
	c = a[1]

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)

	a[1] = value()
	fmt.Println("change a", a)
}

func value() [9]int {
	var retvalue = [9]int{726, 726, 726, 726, 726, 726, 726, 726, 726}
	return retvalue
}
