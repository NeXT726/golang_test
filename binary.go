package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type student struct {
	age uint32
	score uint32
}

func main(){
	var dujianglin = student{21, 80}
	fmt.Println(dujianglin)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, dujianglin)
	if err!= nil {
		fmt.Println("wrong:", err)
	}
	var s = buf.Bytes()
	var prt = student{binary.LittleEndian.Uint32(s), binary.LittleEndian.Uint32(s[4:])}
	fmt.Println(prt)
}
