package main

import (
	"fmt"

	"github.com/skkim-01/wake-on-lan/wol"
)

var mac1 = "484520396e11"
var mac2 = "705DCCF9136B"

func main() {
	fmt.Println("#INFO\tWOL Tester")

	//wol.GetPacket("aabbccddeeff")

	byteMagicPacket := wol.NewPacket(mac1)

	wol.Wake(byteMagicPacket, "255.255.255.255:7")
	wol.Wake(byteMagicPacket, "255.255.255.255:8")
	wol.Wake(byteMagicPacket, "255.255.255.255:9")

	mp2 := wol.NewPacket(mac2)
	wol.Wake(mp2, "255.255.255.255:7")
	wol.Wake(mp2, "255.255.255.255:8")
	wol.Wake(mp2, "255.255.255.255:9")
}

//  32F6#600002
