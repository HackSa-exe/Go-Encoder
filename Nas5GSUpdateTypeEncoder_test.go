package Nas5GSUpdateTypeEncoder

import (
	"bytes"
	"testing"
)

// Checks the struct "nasnas5GSUpdateType" and the function "encode" with numeric values
func TestNas5GSUpdateType(t *testing.T) {

	var testStruct nas5GSUpdateType
	testStruct.updateTypeIEI = 1
	testStruct.updateTypeLength = 2
	testStruct.EPS_PNB_CIoT = 0
	testStruct._5GS_PNB_CIoT = 0
	testStruct.ng_RAN_RCU = 1
	testStruct.smsRequested = 1

	var buffer bytes.Buffer
	bufferPtr := &buffer

	testStruct.Encode(bufferPtr)

}
