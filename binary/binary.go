package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type student struct {
	age uint8
	score uint8
}

func main(){
	var dujianglin = student{21, 80}
	fmt.Println(dujianglin)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, dujianglin)
	if err!= nil {
		fmt.Println("wrong:", err)
	}
	fmt.Printf("%d", buf.Bytes())
}
