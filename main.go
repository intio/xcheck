// This code is in public domain.

package main

import (
	"log"

	"github.com/BurntSushi/xgb"
)

func main() {
	xc, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	if xc == nil {
		log.Fatal("connection is nil")
	}
}
