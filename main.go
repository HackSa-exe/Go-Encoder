package main

import (
	"bytes"
	"fmt"

	"github.com/HewlettPackard/structex"
)

type nas5GSUpdateType struct {

	//first octet
	updateTypeIEI byte

	//second octet
	updateTypeLength byte

	//third octet
	sMSrequested                     uint8 `bitfield:"1,"`
	nG_RAN_RCU                       uint8 `bitfield:"1,"`
	preferred_CIoT_Network_Behaviour uint8 `bitfield:"2,"`
	ePS_PNB_CIoT                     uint8 `bitfield:"2,"`
	spare1                           uint8 `bitfield:"1,reserved"`
	spare2                           uint8 `bitfield:"1,reserved"`
}

func main() {

}

func (ie nas5GSUpdateType) Encode(buffer *bytes.Buffer) {

	byteArray, err := structex.EncodeByteBuffer(ie)

	if err != nil {
		panic("Encoding not possible.")
	}

	for index, element := range byteArray {
		buffer.WriteByte(element)
		fmt.Println("Byte index ", index, " Element: ", element)
	}

	fmt.Println(buffer)
}
