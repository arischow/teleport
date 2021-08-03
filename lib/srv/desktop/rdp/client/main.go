package main

/*
#cgo LDFLAGS: -L${SRCDIR}/target/debug -l:librdp_client.a -lpthread -lcrypto -ldl -lssl -lm
#include <librdprs.h>
*/
import "C"
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go start")
	if len(os.Args) < 4 {
		fmt.Println("usage:", os.Args[0], "host:port user password")
		os.Exit(1)
	}
	addr := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	C.connect_rdp(
		cgoString(addr),
		cgoString(username),
		cgoString(password),
		C.uint16_t(800),
		C.uint16_t(600),
	)
	fmt.Println("Go end")
}

func cgoString(s string) C.CGOString {
	sb := []byte(s)
	return C.CGOString{
		data: (*C.uint8_t)(C.CBytes(sb)),
		len:  C.uint16_t(len(sb)),
	}
}
