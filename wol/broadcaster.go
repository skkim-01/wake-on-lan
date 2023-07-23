package wol

import (
	"fmt"
	"net"
)

func Wake(mp []byte, bcastAddr string) {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	broadcastAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		panic(err)
	}

	udpConn, err := net.DialUDP("udp", nil, broadcastAddr)
	if err != nil {
		panic(err)
	}

	nBytesSent, err := udpConn.Write(mp)
	fmt.Println("#INFO\t", nBytesSent, "sent / ERR:", err)
}
