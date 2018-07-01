package main

import (
	"DavidChain/basicPrototype14/BLC"
)

func main() {


	// 创世区块
	blockchain := BLC.CreateGenesisBlockChain()
	defer blockchain.DB.Close()
	cli := BLC.CLI{blockchain}
	cli.Run()

}
