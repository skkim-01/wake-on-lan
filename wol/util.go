package wol

import (
	"encoding/hex"
	"fmt"
)

func debugBytes(raw []byte) {
	fmt.Println("#INFO\tdebugBytes:", len(raw))

	fmt.Println("    00 01 02 03 04 05 06 07")
	fmt.Println("===========================")
	fmt.Printf("00 |")
	for i := 1; i < len(raw)+1; i++ {
		fmt.Print(byte2hex(raw[i-1:i], false), " ")
		if i != 0 && i%8 == 0 {
			fmt.Println()
			fmt.Printf("%02X |", i/8)
		}
	}
	fmt.Println()
	fmt.Println()
}

func memcpy(src []byte, target []byte, size int) {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	for i := 0; i < size; i++ {
		src[i] = target[i]
	}
}

func byte2hex(b []byte, bAppender bool) string {
	retv := hex.EncodeToString(b)

	if bAppender {
		if 0 == len(retv)%2 {
			retv = "0x" + retv
		} else {
			retv = "0x0" + retv
		}
	}
	return retv
}
