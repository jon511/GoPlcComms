package Controller

import (
	"net"
)

type Controller struct{

	ipAddress string
	processorSlot int
	micro800 bool
	port int

	vendorID []byte
	context []byte
	contextPointer int

	conn net.Conn
	setTimeout float64
	isConnected bool
	otNetworkConnectionID []byte
	sessionHandle []byte
	sessionIsRegistered bool
	serialNumber []byte
	sequenceCounter int
	offset int
	structIdentifier []byte

	initialized bool

	knownTags []string
}
