package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type student struct {
	Age uint32
	Score uint32
}

func main(){
	buf := new(bytes.Buffer)

	var dujianglin = student{21, 80}
	fmt.Println("dujianglin", dujianglin)
	err := binary.Write(buf, binary.LittleEndian, dujianglin)
	if err!= nil {
		fmt.Println("wrong:", err)
	}

	var djl = student{1, 1}
	err = binary.Read(buf, binary.LittleEndian, &djl)
	if err!= nil {
		fmt.Println("wrong:", err)
	}	
	fmt.Println("djl", djl)
}
