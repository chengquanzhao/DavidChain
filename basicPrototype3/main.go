package main

import (
	"DavidChain/basicPrototype3/BLC"
	"fmt"
)

func main() {
	blockchain := BLC.CreateGenesisBlockChain()
	fmt.Println(blockchain)

	fmt.Println(blockchain.Blocks)
	fmt.Println(blockchain.Blocks[0])

}
