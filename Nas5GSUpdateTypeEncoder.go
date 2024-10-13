package Nas5GSUpdateTypeEncoder

import (
	"bytes"
	"fmt"

	"github.com/HewlettPackard/structex"
)

//	My take on the task i received from CampusGenius.
//	Full Task can be found in the subfolder \Additional Information\CampusGenius_GO-Testtask.pdf.

// Defines a struct with three bytes(octets), in which the third one is written in Structex convention for encoding.
// Specifications for this can be found in the subfolder \Additional Information\Technical Specifications.docx.
type nas5GSUpdateType struct {

	//first octet
	updateTypeIEI byte

	//second octet
	updateTypeLength byte

	//third octet
	smsRequested  uint8 `bitfield:"1,"`
	ng_RAN_RCU    uint8 `bitfield:"1,"`
	_5GS_PNB_CIoT uint8 `bitfield:"2,"`
	EPS_PNB_CIoT  uint8 `bitfield:"2,"`
	spare1        uint8 `bitfield:"1,reserved"`
	spare2        uint8 `bitfield:"1,reserved"`
}

// Uses Structex to encode an nas5GSUpdateType struct into an buffer
func (ie nas5GSUpdateType) Encode(buffer *bytes.Buffer) {

	//Structex func returns byteArray and and error for further use
	byteArray, err := structex.EncodeByteBuffer(ie)
	if err != nil {
		panic("Encoding not possible.")
	}

	//iterates through the byteArray and writes them into the buffer
	for index, element := range byteArray {
		buffer.WriteByte(element)
		fmt.Println("Byte index ", index, " Element: ", element)
	}

	fmt.Println(buffer)
}
