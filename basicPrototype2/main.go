package main

import (
	"DavidChain/basicPrototype2/BLC"
	"fmt"
)

func main() {
	block := BLC.CreateGenesisBlock("Genesis block")
	fmt.Println(block)
}
