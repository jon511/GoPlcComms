package EIP

import "fmt"

func (c *Controller) BuildEipHeader(tagIOI []byte) []byte {

	if c.contextPointer == 155 {
		c.contextPointer = 0
	}

	fmt.Println("tagIOI")
	fmt.Println(tagIOI)

	//eipPayLoadLength := byte(22 + len(tagIOI))
	eipConnectedDataLength := byte(len(tagIOI) + 2)
	fmt.Printf("dl %d", eipConnectedDataLength)
	eipCommand := []byte{0x70, 0x00}

	temp := int32ToSliceOfBytes(true, 22 + len(tagIOI), 2)
	fmt.Println(len(tagIOI))
	fmt.Println("eip length")
	printHex(temp)
	eipLength := temp
	eipSessionHandle := c.sessionHandle
	eipStatus := []byte {0x00, 0x00, 0x00, 0x00}

	eipContext := contextMap[c.contextPointer]
	c.contextPointer ++
	fmt.Printf("context header: %d", c.contextPointer)
	eipOptions := []byte{0x00, 0x00, 0x00, 0x00}
	eipInterfaceHandle := []byte{0x00, 0x00, 0x00, 0x00}
	eipTimeout := []byte{0x00, 0x00}
	eipItemCount := []byte{0x02, 0x00}
	eipItem1ID := []byte{0xa1, 0x00}
	eipItem1Length := []byte{0x04, 0x00}
	eipItem1 := c.otNetWorkConnectionID
	fmt.Println("otNetworkID")
	printHex(c.otNetWorkConnectionID)
	eipItem2ID := []byte{0xb1, 0x00}
	eipItem2Length := []byte{byte(eipConnectedDataLength), 0x00}
	printHex(eipItem2Length)
	eipSequence := int32ToSliceOfBytes(true, c.sequenceCounter, 2)
	c.sequenceCounter ++
	c.sequenceCounter = c.sequenceCounter % 0x10000

	var eipHeaderFrame []byte
	eipHeaderFrame = append(eipHeaderFrame, eipCommand...)
	eipHeaderFrame = append(eipHeaderFrame, eipLength...)
	eipHeaderFrame = append(eipHeaderFrame, eipSessionHandle...)
	eipHeaderFrame = append(eipHeaderFrame, eipStatus...)
	eipHeaderFrame = append(eipHeaderFrame, eipContext...)
	eipHeaderFrame = append(eipHeaderFrame, eipOptions...)
	eipHeaderFrame = append(eipHeaderFrame, eipInterfaceHandle...)
	eipHeaderFrame = append(eipHeaderFrame, eipTimeout...)
	eipHeaderFrame = append(eipHeaderFrame, eipItemCount...)
	eipHeaderFrame = append(eipHeaderFrame, eipItem1ID...)
	eipHeaderFrame = append(eipHeaderFrame, eipItem1Length...)
	eipHeaderFrame = append(eipHeaderFrame, eipItem1...)
	eipHeaderFrame = append(eipHeaderFrame, eipItem2ID...)
	eipHeaderFrame = append(eipHeaderFrame, eipItem2Length...)
	eipHeaderFrame = append(eipHeaderFrame, eipSequence...)
	fmt.Println("eip header")
	printHex(eipHeaderFrame)
	return append(eipHeaderFrame, tagIOI...)

}