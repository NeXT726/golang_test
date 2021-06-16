//尝试已有的一个rdmamap包。只有设备的get等操作，用途其实不是很大。

package main

import (
	"fmt"

	"github.com/Mellanox/rdmamap"
)

func main() {
	rdmaDevices := rdmamap.GetRdmaDeviceList()
	fmt.Println("Devices: ", rdmaDevices)

	for _, dev := range rdmaDevices {
		charDevices := rdmamap.GetRdmaCharDevices(dev)
		fmt.Printf("Rdma device: = %s", dev)
		fmt.Println(" Char devices: = ", charDevices)
	}
}
