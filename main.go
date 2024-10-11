package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/HewlettPackard/structex"
)

type bitfield struct {
	sMSrequested                     uint8 `bitfield:"1,"`
	nG_RAN_RCU                       uint8 `bitfield:"1,"`
	preferred_CIoT_Network_Behaviour uint8 `bitfield:"2,"`
	ePS_PNB_CIoT                     uint8 `bitfield:"2,"`
	spare1                           uint8 `bitfield:"1,reserved"`
	spare2                           uint8 `bitfield:"1,reserved"`
}

type nas5GSUpdateType struct {

	//first octet
	updateTypeIEI byte

	//second octet
	updateTypeLength byte

	//third octet
	bitfield
}

func main() {

	var test nas5GSUpdateType
	test.updateTypeIEI = 1
	test.updateTypeLength = 2
	test.bitfield.ePS_PNB_CIoT = 0
	test.bitfield.preferred_CIoT_Network_Behaviour = 0
	test.bitfield.ePS_PNB_CIoT = 1
	test.bitfield.sMSrequested = 1

	var buffer bytes.Buffer
	bufferPtr := &buffer

	Encode(test, bufferPtr)

}

func Encode(ie nas5GSUpdateType, buffer *bytes.Buffer) {

	//write first octet in buffer
	fmt.Println("Update Type: ", ie.updateTypeIEI)
	buffer.WriteByte(ie.updateTypeIEI)

	//write second octet in buffer
	fmt.Println("Update Type length: ", ie.updateTypeLength)
	buffer.WriteByte(ie.updateTypeLength)

	fmt.Println("Buffered Stuff ", buffer.String())

	//convert bitfield and add to buffer
	var writer io.ByteWriter
	structex.Encode(writer, ie.bitfield)
	fmt.Println(writer)
}
