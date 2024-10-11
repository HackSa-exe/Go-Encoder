package main

import (
	"bytes"
	"testing"
)

func TestAdd(t *testing.T) {

	var test nas5GSUpdateType
	test.updateTypeIEI = 1
	test.updateTypeLength = 2
	test.ePS_PNB_CIoT = 0
	test.preferred_CIoT_Network_Behaviour = 0
	test.nG_RAN_RCU = 1
	test.sMSrequested = 1

	var buffer bytes.Buffer
	bufferPtr := &buffer

	test.Encode(bufferPtr)

}
