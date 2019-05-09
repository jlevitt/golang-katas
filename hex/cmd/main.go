package main

import (
	"fmt"
	"github.com/jlevitt/katas/hex"
)

func main() {
	fmt.Printf("'%v'\n", hex.HexToBin("11"))
	fmt.Printf("'%v'\n", hex.BinToHex("11001101"))
}
