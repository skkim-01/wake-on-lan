package wol

const PACKET_SIZE = 102

func NewPacket(macAddr string) []byte {
	retvBytes := make([]byte, 102, 102)

	// #1. Fill 0xFF: 0~6
	for i := 0; i < 6; i++ {
		retvBytes[i] = 0xff
	}

	// #2. get MacAddrByte
	macBuffer := _macAddrBuffer(macAddr)

	// #3. fillPacket
	_fillPacket( retvBytes, macBuffer )

	// #optional. debugprint
	debugBytes(retvBytes)

	return retvBytes
}

func _macAddrBuffer(macAddr string) []byte {
	byteRetv := make([]byte, 6, 6)
	buffer := make([]byte, 3, 3)

	for i :=0 ; i < 6 ; i++ {
		buffer = []byte( macAddr[ i*2 : i*2+2] )
		byteRetv[i] = buffer[0]
	}

	return byteRetv
}

func _fillPacket( src []byte, macBuffer []byte ) {	

	for i := 0 ; i < 16 ; i++ {
		memcpy( src[(i+1)*6:], macBuffer, len(macBuffer) )
	}
}